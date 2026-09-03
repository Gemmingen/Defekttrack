package main

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"defekttrack/api"
)

type Server struct {
	DB *sql.DB
}

// GET /laptops (mit optionalem Filter nach Fehlerkategorie)
func (s *Server) GetLaptops(ctx echo.Context, params api.GetLaptopsParams) error {
	var rows *sql.Rows
	var err error

	if params.Fehler != nil && string(*params.Fehler) != "" {
		query := "SELECT id, marke, name, os, fehler FROM laptops WHERE fehler = $1"
		rows, err = s.DB.Query(query, string(*params.Fehler))
	} else {
		query := "SELECT id, marke, name, os, fehler FROM laptops"
		rows, err = s.DB.Query(query)
	}

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "DB-Fehler: "+err.Error())
	}
	defer rows.Close()

	laptops := []api.Laptop{}
	for rows.Next() {
		var l api.Laptop
		var fehlerStr string
		if err := rows.Scan(&l.Id, &l.Marke, &l.Name, &l.Os, &fehlerStr); err != nil {
			return ctx.String(http.StatusInternalServerError, "Parsing-Fehler: "+err.Error())
		}
		l.Fehler = api.FehlerKategorie(fehlerStr)
		laptops = append(laptops, l)
	}

	return ctx.JSON(http.StatusOK, laptops)
}

// POST /laptops
func (s *Server) PostLaptops(ctx echo.Context) error {
	var input api.LaptopInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON-Format")
	}

	var newID int
	query := `INSERT INTO laptops (marke, name, os, fehler) VALUES ($1, $2, $3, $4) RETURNING id`
	err := s.DB.QueryRow(query, input.Marke, input.Name, input.Os, string(input.Fehler)).Scan(&newID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Einfügefehler: "+err.Error())
	}

	return ctx.JSON(http.StatusCreated, api.Laptop{
		Id:     newID,
		Marke:  input.Marke,
		Name:   input.Name,
		Os:     input.Os,
		Fehler: input.Fehler,
	})
}

// GET /laptops/{id}
func (s *Server) GetLaptopsId(ctx echo.Context, id int) error {
	var l api.Laptop
	var fehlerStr string

	query := "SELECT id, marke, name, os, fehler FROM laptops WHERE id = $1"
	err := s.DB.QueryRow(query, id).Scan(&l.Id, &l.Marke, &l.Name, &l.Os, &fehlerStr)
	if errors.Is(err, sql.ErrNoRows) {
		return ctx.String(http.StatusNotFound, "Laptop nicht gefunden")
	} else if err != nil {
		return ctx.String(http.StatusInternalServerError, "DB-Fehler: "+err.Error())
	}
	l.Fehler = api.FehlerKategorie(fehlerStr)

	// Zubehörige Logs laden
	logRows, err := s.DB.Query("SELECT id, bearbeiter, notiz, timestamp FROM logs WHERE laptop_id = $1 ORDER BY timestamp DESC", id)
	if err == nil {
		defer logRows.Close()
		logs := []api.LogEintrag{}
		for logRows.Next() {
			var logEntry api.LogEintrag
			if err := logRows.Scan(&logEntry.Id, &logEntry.Bearbeiter, &logEntry.Notiz, &logEntry.Timestamp); err == nil {
				logs = append(logs, logEntry)
			}
		}
		l.ItLogs = &logs
	}

	return ctx.JSON(http.StatusOK, l)
}

// PUT /laptops/{id}
func (s *Server) PutLaptopsId(ctx echo.Context, id int) error {
	var input api.LaptopInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON-Format")
	}

	query := "UPDATE laptops SET marke = $1, name = $2, os = $3, fehler = $4 WHERE id = $5"
	res, err := s.DB.Exec(query, input.Marke, input.Name, input.Os, string(input.Fehler), id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Update-Fehler: "+err.Error())
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ctx.String(http.StatusNotFound, "Laptop nicht gefunden")
	}

	return ctx.JSON(http.StatusOK, api.Laptop{
		Id:     id,
		Marke:  input.Marke,
		Name:   input.Name,
		Os:     input.Os,
		Fehler: input.Fehler,
	})
}

// DELETE /laptops/{id}
func (s *Server) DeleteLaptopsId(ctx echo.Context, id int) error {
	res, err := s.DB.Exec("DELETE FROM laptops WHERE id = $1", id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Löschfehler: "+err.Error())
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ctx.String(http.StatusNotFound, "Laptop nicht gefunden")
	}

	return ctx.NoContent(http.StatusNoContent)
}

// POST /laptops/{id}/logs
func (s *Server) PostLaptopsIdLogs(ctx echo.Context, id int) error {
	var input api.LogEintragInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON-Format")
	}

	var newID int
	var ts time.Time
	query := "INSERT INTO logs (laptop_id, bearbeiter, notiz) VALUES ($1, $2, $3) RETURNING id, timestamp"
	err := s.DB.QueryRow(query, id, input.Bearbeiter, input.Notiz).Scan(&newID, &ts)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Fehler beim Erstellen des Log-Eintrags: "+err.Error())
	}

	return ctx.JSON(http.StatusCreated, api.LogEintrag{
		Id:         newID,
		Bearbeiter: input.Bearbeiter,
		Notiz:      input.Notiz,
		Timestamp:  ts,
	})
}

// PUT /laptops/{id}/logs/{logId}
func (s *Server) PutLaptopsIdLogsLogId(ctx echo.Context, id int, logId int) error {
	var input api.LogEintragInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON-Format")
	}

	var ts time.Time
	query := "UPDATE logs SET bearbeiter = $1, notiz = $2 WHERE id = $3 AND laptop_id = $4 RETURNING timestamp"
	err := s.DB.QueryRow(query, input.Bearbeiter, input.Notiz, logId, id).Scan(&ts)
	if errors.Is(err, sql.ErrNoRows) {
		return ctx.String(http.StatusNotFound, "Log-Eintrag nicht gefunden")
	} else if err != nil {
		return ctx.String(http.StatusInternalServerError, "Update-Fehler: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, api.LogEintrag{
		Id:         logId,
		Bearbeiter: input.Bearbeiter,
		Notiz:      input.Notiz,
		Timestamp:  ts,
	})
}

// DELETE /laptops/{id}/logs/{logId}
func (s *Server) DeleteLaptopsIdLogsLogId(ctx echo.Context, id int, logId int) error {
	res, err := s.DB.Exec("DELETE FROM logs WHERE id = $1 AND laptop_id = $2", logId, id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Löschfehler: "+err.Error())
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ctx.String(http.StatusNotFound, "Log-Eintrag nicht gefunden")
	}

	return ctx.NoContent(http.StatusNoContent)
}
