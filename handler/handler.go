package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"defekttrack/api"
	"defekttrack/service"
)

type Server struct {
	Service service.LaptopService
}

func New(svc service.LaptopService) api.ServerInterface {
	return &Server{Service: svc}
}

func (s *Server) GetLaptops(ctx echo.Context, params api.GetLaptopsParams) error {
	filter := ""
	if params.Fehler != nil {
		filter = string(*params.Fehler)
	}
	laptops, err := s.Service.GetLaptops(filter)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, laptops)
}

func (s *Server) PostLaptops(ctx echo.Context) error {
	var input api.LaptopInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON")
	}
	laptop, err := s.Service.CreateLaptop(input)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, laptop)
}

func (s *Server) GetLaptopsId(ctx echo.Context, id int) error {
	laptop, err := s.Service.GetLaptopByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.String(http.StatusNotFound, "Laptop nicht gefunden")
	} else if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, laptop)
}

func (s *Server) PutLaptopsId(ctx echo.Context, id int) error {
	var input api.LaptopInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON")
	}
	laptop, err := s.Service.UpdateLaptop(id, input)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.String(http.StatusNotFound, "Laptop nicht gefunden")
	} else if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, laptop)
}

func (s *Server) DeleteLaptopsId(ctx echo.Context, id int) error {
	err := s.Service.DeleteLaptop(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.String(http.StatusNotFound, "Laptop nicht gefunden")
	} else if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) PostLaptopsIdLogs(ctx echo.Context, id int) error {
	var input api.LogEintragInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON")
	}
	logEntry, err := s.Service.CreateLog(id, input)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, logEntry)
}

func (s *Server) PutLaptopsIdLogsLogId(ctx echo.Context, id int, logId int) error {
	var input api.LogEintragInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, "Ungültiges JSON")
	}
	logEntry, err := s.Service.UpdateLog(id, logId, input)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.String(http.StatusNotFound, "Log nicht gefunden")
	} else if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, logEntry)
}

func (s *Server) DeleteLaptopsIdLogsLogId(ctx echo.Context, id int, logId int) error {
	err := s.Service.DeleteLog(id, logId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.String(http.StatusNotFound, "Log nicht gefunden")
	} else if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
