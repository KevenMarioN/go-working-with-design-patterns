package main

import (
	"database/sql"
	"flag"
	"fmt"
	"gobreeders/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

const PORT = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      configApp
	db          *sql.DB
	Models      models.Models
}

type configApp struct {
	port     string
	useCache bool
	dsn      string
}

func main() {
	app := application{
		config: configApp{
			port: PORT,
		},
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.StringVar(&app.config.port, "port", PORT, fmt.Sprintf("Use port %s", app.config.port))
	flag.Parse()

	var err error
	if app.db, err = initMySQLB(app.config.dsn); err != nil {
		log.Panic(err)
	}
	app.Models = *models.New(app.db)

	srv := &http.Server{
		Addr:              app.config.port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	log.Printf("Starting web application on port %s, cache: %t", PORT, app.config.useCache)
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
