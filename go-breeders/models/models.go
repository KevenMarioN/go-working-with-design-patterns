package models

import "time"

type DogBreed struct {
	Breed
}

type CatBreed struct {
	Breed
}

type Breed struct {
	ID             int    `json:"id"`
	Breed          string `json:"breed"`
	WeightLowLbs   int    `json:"weight_low_lbs"`
	WeightHighLbs  int    `json:"weight_high_lbs"`
	AverageWeight  int    `json:"average_weight"`
	Lifespan       int    `json:"lifespan"`
	Details        string `json:"details"`
	AlternateNames string `json:"alternate_names"`
	GeographOrigin string `json:"geograph_origin"`
}

type Dog struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateBirth        time.Time `json:"date_birth"`
	SpayedOrNeutered int       `json:"spayed_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Cat struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateBirth        time.Time `json:"date_birth"`
	SpayedOrNeutered int       `json:"spayed_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Breeder struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Address   string      `json:"address"`
	City      string      `json:"city"`
	ProvState string      `json:"prov_state"`
	Country   string      `json:"country"`
	Zip       string      `json:"zip"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	Active    int         `json:"active"`
	DogBreeds []*DogBreed `json:"dog_breeds"`
	CatBreeds []*CatBreed `json:"cat_breeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	LifeSpan    int    `json:"life_span"`
}
