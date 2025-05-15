package pets

import (
	"fmt"
	"gobreeders/models"
)

type AnimalSpecies string

const (
	DOG AnimalSpecies = "dog"
	CAT AnimalSpecies = "cat"
)

type IAnimal interface {
	Show() string
}

type DogFactory struct {
	Pet *models.Dog
}

func (df *DogFactory) Show() string {
	return fmt.Sprintf("This animal is a %v", df.Pet.Breed.Breed)
}

type CatFactory struct {
	pet *models.Cat
}

func (ct *CatFactory) Show() string {
	return fmt.Sprintf("This animal is a %v", ct.pet.Breed.Breed)
}

type IPetFactory interface {
	newPet() IAnimal
}

type DogAbastractFactory struct{}

func (df *DogAbastractFactory) newPet() IAnimal {
	return &DogFactory{
		Pet: &models.Dog{},
	}
}

type CatAbastractFactory struct{}

func (df *CatAbastractFactory) newPet() IAnimal {
	return &CatFactory{
		pet: &models.Cat{},
	}
}

func NewPetAbstractFactory(species string) (IAnimal, error) {
	switch AnimalSpecies(species) {
	case DOG:
		daf := DogAbastractFactory{}
		dog := daf.newPet()
		return dog, nil
	case CAT:
		caf := CatAbastractFactory{}
		cat := caf.newPet()
		return cat, nil
	default:
		return nil, fmt.Errorf("error this pet not exist %s", species)
	}
}
