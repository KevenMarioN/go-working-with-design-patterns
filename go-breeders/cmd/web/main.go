package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"
)

const PORT = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      configApp
}

type configApp struct {
	port     string
	useCache bool
}

func main() {
	app := application{
		config: configApp{
			port: PORT,
		},
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template chace")
	flag.Parse()

	srv := &http.Server{
		Addr:              app.config.port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	log.Printf("Starting web application on port %s, cache: %t", PORT, app.config.useCache)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
