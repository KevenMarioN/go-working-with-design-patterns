package main

import (
	"encoding/json"
	"fmt"
	"gobreeders/pets"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
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

func (a *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var tool toolbox.Tools
	_ = tool.WriteJSON(w, http.StatusCreated, pets.NewPet("dog"))
}

func (a *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var tool toolbox.Tools
	_ = tool.WriteJSON(w, http.StatusCreated, pets.NewPet("cat"))
}

func (a *application) CreateDogOrCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var tool toolbox.Tools
	species := chi.URLParam(r, "species")
	pet, err := pets.NewPetAbstractFactory(species)
	if err != nil {
		_ = tool.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = tool.WriteJSON(w, http.StatusCreated, pet)
}

func (a *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var tool toolbox.Tools
	dogBreeds, err := a.Models.DogBreed.All()
	if err != nil {
		_ = tool.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = tool.WriteJSON(w, http.StatusCreated, dogBreeds)
}

func (a *application) GetAllCatBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var tool toolbox.Tools
	dogBreeds, err := a.Models.CatBreed.All()
	if err != nil {
		_ = tool.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = tool.WriteJSON(w, http.StatusCreated, dogBreeds)
}

func (a *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	a.render(w, "test.page.gohtml", nil)
}
