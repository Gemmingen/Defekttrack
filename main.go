package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"defekttrack/api"
	"defekttrack/handler"
	"defekttrack/repository"
	"defekttrack/service"
)

const connStr = "host=localhost port=5432 user=postgres password=secret dbname=laptopdb sslmode=disable"

func main() {
	db, err := repository.ConnectAndMigrate(connStr)
	if err != nil {
		log.Fatalf("Datenbank-Fehler: %v", err)
	}
	fmt.Println("Datenbank verbunden & migriert!")

	repo := repository.New(db)

	svc := service.New(repo)

	srv := handler.New(svc)

	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Printf("HTTP %s | Status: %d | URI: %s", v.Method, v.Status, v.URI)
			return nil
		},
	}))
	api.RegisterHandlers(e, srv)

	// Swagger UI
	e.Static("/openapi.yaml", "openapi.yaml")
	e.GET("/swagger", func(c echo.Context) error {
		html := `<!DOCTYPE html><html><head><link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" /></head><body><div id="swagger-ui"></div><script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script><script>window.onload = () => { window.ui = SwaggerUIBundle({ url: '/openapi.yaml', dom_id: '#swagger-ui' }); };</script></body></html>`
		return c.HTML(200, html)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
