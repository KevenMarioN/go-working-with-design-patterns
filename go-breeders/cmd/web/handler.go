package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) Pong(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
}

func (a *application) Home(w http.ResponseWriter, r *http.Request) {
	a.render(w, "home.page.gohtml", nil)
}

func (a *application) SelectPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	a.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}
