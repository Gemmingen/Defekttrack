package main

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"defekttrack/api"
	"defekttrack/handler"
	"defekttrack/repository"
	"defekttrack/service"
)

const connStr = "host=localhost port=5432 user=postgres password=secret dbname=laptopdb sslmode=disable"

func main() {
	fx.New(
		// 1. Abhängigkeiten registrieren (wie builder.Services in C#)
		fx.Provide(
			func() (*gorm.DB, error) {
				db, err := repository.ConnectAndMigrate(connStr)
				if err == nil {
					fmt.Println("Datenbank verbunden & migriert!")
				}
				return db, err
			},
			repository.New, // Baut repository.Repository
			service.New,    // Baut service.LaptopService
			handler.New,    // Baut api.ServerInterface
			echo.New,       // Baut *echo.Echo
		),
		// 2. Anwendung starten
		fx.Invoke(registerRoutesAndStart),
	).Run()
}

func registerRoutesAndStart(lc fx.Lifecycle, e *echo.Echo, srv api.ServerInterface) {
	// Logger Middleware
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Printf("HTTP %s | Status: %d | URI: %s", v.Method, v.Status, v.URI)
			return nil
		},
	}))

	// API Handlers
	api.RegisterHandlers(e, srv)

	// Swagger UI (HTML-String bereinigt)
	e.Static("/openapi.yaml", "openapi.yaml")
	e.GET("/swagger", func(c echo.Context) error {
		html := `<!DOCTYPE html><html><head><link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" /></head><body><div id="swagger-ui"></div><script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script><script>window.onload = () => { window.ui = SwaggerUIBundle({ url: '/openapi.yaml', dom_id: '#swagger-ui' }); };</script></body></html>`
		return c.HTML(200, html)
	})

	// Server Start / Shutdown Hook
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Start(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
