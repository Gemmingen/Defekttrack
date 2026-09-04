package service

import (
	"database/sql"
	"defekttrack/api"
	"defekttrack/repository"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=secret dbname=laptopdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("DB-Verbindungsfehler: %v", err)
	}
	defer db.Close()

	// Prüfen ob DB erreichbar
	if err := db.Ping(); err != nil {
		log.Fatalf("DB nicht erreichbar: %v", err)
	}
	fmt.Println("Erfolgreich mit Postgres verbunden!")

	createTableSQL := `
CREATE TABLE IF NOT EXISTS laptops (
    id SERIAL PRIMARY KEY,
    marke TEXT NOT NULL,
    name TEXT NOT NULL,
    os TEXT NOT NULL,
    fehler TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    laptop_id INT NOT NULL REFERENCES laptops(id) ON DELETE CASCADE,
    bearbeiter TEXT NOT NULL,
    notiz TEXT NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW()
);`
	//Tabellen erstellen
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatalf("Fehler beim Erstellen der Tabelle: %v", err)
	}

	//Echo und Backend
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

	meinBackend := &repository.Server{DB: db}
	api.RegisterHandlers(e, meinBackend)

	//Swagger
	e.Static("/openapi.yaml", "openapi.yaml")
	e.GET("/swagger", func(c echo.Context) error {
		html := `<!DOCTYPE html>
		<html>
		<head>
			<link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
		</head>
		<body>
			<div id="swagger-ui"></div>
			<script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
			<script>
				window.onload = () => { window.ui = SwaggerUIBundle({ url: '/openapi.yaml', dom_id: '#swagger-ui' }); };
			</script>
		</body>
		</html>`
		return c.HTML(200, html)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
