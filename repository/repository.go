package repository

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"defekttrack/api"
	"defekttrack/repository/models"
)

type Server struct {
	DB *gorm.DB
}

// GET /laptops
// GET /laptops
func (s *Server) GetLaptops(ctx echo.Context, params api.GetLaptopsParams) error {
	var laptopModels []models.LaptopModel

	// Preload lädt die Logs aller Laptops in einem Aufruf mit
	query := s.DB.Preload("Logs")

	if params.Fehler != nil && string(*params.Fehler) != "" {
		query = query.Where("fehler = ?", string(*params.Fehler))
	}

	if err := query.Find(&laptopModels).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "DB-Fehler: "+err.Error())
	}

	laptops := make([]api.Laptop, 0, len(laptopModels))
	for _, m := range laptopModels {
		laptops = append(laptops, m.ToAPI())
	}

	return ctx.JSON(http.StatusOK, laptops)
}

// POST /laptops
func (s *Server) PostLaptops(ctx echo.Context) error {
	var input api.LaptopInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON")
	}

	model := models.ToLaptopModel(input)
	if err := s.DB.Create(&model).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Fehler beim Speichern: "+err.Error())
	}

	return ctx.JSON(http.StatusCreated, model.ToAPI())
}

// GET /laptops/{id}
func (s *Server) GetLaptopsId(ctx echo.Context, id int) error {
	var model models.LaptopModel
	if err := s.DB.Preload("Logs").First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.String(http.StatusNotFound, "Laptop nicht gefunden")
		}
		return ctx.String(http.StatusInternalServerError, "DB-Fehler: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, model.ToAPI())
}

// PUT /laptops/{id}
func (s *Server) PutLaptopsId(ctx echo.Context, id int) error {
	var input api.LaptopInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON")
	}

	var model models.LaptopModel
	if err := s.DB.First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.String(http.StatusNotFound, "Laptop nicht gefunden")
		}
		return ctx.String(http.StatusInternalServerError, "DB-Fehler: "+err.Error())
	}

	model.Marke = input.Marke
	model.Name = input.Name
	model.Os = input.Os
	model.Fehler = string(input.Fehler)

	if err := s.DB.Save(&model).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Update-Fehler: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, model.ToAPI())
}

// DELETE /laptops/{id}
func (s *Server) DeleteLaptopsId(ctx echo.Context, id int) error {
	res := s.DB.Delete(&models.LaptopModel{}, id)
	if res.Error != nil {
		return ctx.String(http.StatusInternalServerError, "Löschfehler: "+res.Error.Error())
	}
	if res.RowsAffected == 0 {
		return ctx.String(http.StatusNotFound, "Laptop nicht gefunden")
	}

	return ctx.NoContent(http.StatusNoContent)
}

// POST /laptops/{id}/logs
func (s *Server) PostLaptopsIdLogs(ctx echo.Context, id int) error {
	var input api.LogEintragInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON")
	}

	logModel := models.ToLogModel(input, id)
	if err := s.DB.Create(&logModel).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Fehler beim Erstellen des Logs: "+err.Error())
	}

	return ctx.JSON(http.StatusCreated, logModel.ToAPI())
}

// PUT /laptops/{id}/logs/{logId}
func (s *Server) PutLaptopsIdLogsLogId(ctx echo.Context, id int, logId int) error {
	var input api.LogEintragInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON")
	}

	var logModel models.LogModel
	if err := s.DB.Where("id = ? AND laptop_id = ?", logId, id).First(&logModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.String(http.StatusNotFound, "Log nicht gefunden")
		}
		return ctx.String(http.StatusInternalServerError, "DB-Fehler: "+err.Error())
	}

	logModel.Bearbeiter = input.Bearbeiter
	logModel.Notiz = input.Notiz

	if err := s.DB.Save(&logModel).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Update-Fehler: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, logModel.ToAPI())
}

// DELETE /laptops/{id}/logs/{logId}
func (s *Server) DeleteLaptopsIdLogsLogId(ctx echo.Context, id int, logId int) error {
	res := s.DB.Where("id = ? AND laptop_id = ?", logId, id).Delete(&models.LogModel{})
	if res.Error != nil {
		return ctx.String(http.StatusInternalServerError, "Löschfehler: "+res.Error.Error())
	}
	if res.RowsAffected == 0 {
		return ctx.String(http.StatusNotFound, "Log nicht gefunden")
	}

	return ctx.NoContent(http.StatusNoContent)
}
