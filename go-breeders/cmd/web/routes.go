package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer, middleware.Logger)
	mux.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", app.Home)
	mux.Get("/{page}", app.SelectPage)
	mux.Get("/test-patterns", app.TestPatterns)

	mux.Route("/api", func(r chi.Router) {
		r.Use(middleware.AllowContentType("Json"))
		r.Get("/ping", app.Pong)
		r.Get("/dog-factory", app.CreateCatFromFactory)
		r.Get("/cat-factory", app.CreateDogFromFactory)
		r.Get("/dog-cat-abstract-factory/{species}", app.CreateDogOrCatFromAbstractFactory)
		r.Get("/dog-breeds", app.GetAllDogBreedsJSON)
		r.Get("/cat-breeds", app.GetAllCatBreedsJSON)
	})

	return mux
}
