package models

import "time"

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightHightLbs   int    `json:"weight_high_lbs"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	AverageWeight    int    `json:"average_weight"`
	LifeSpan         int    `json:"life_span"`
	Details          string `json:"details"`
	AlternativeNames string `json:"alternative_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type CatBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightHightLbs   int    `json:"weight_high_lbs"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	AverageWeight    int    `json:"average_weight"`
	LifeSpan         int    `json:"life_span"`
	Details          string `json:"details"`
	AlternativeNames string `json:"alternative_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type Dog struct {
	ID               int       `json:"id"`
	DogName          string    `json:"dog_name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered bool      `json:"spayed_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Cat struct {
	ID               int       `json:"id"`
	CatName          string    `json:"cat_name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered bool      `json:"spayed_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Breeder struct {
	ID          int         `json:"id"`
	BreederName string      `json:"breeder_name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	ProvState   string      `json:"prov_state"`
	Country     string      `json:"country"`
	Zip         string      `json:"zip"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Active      int         `json:"active"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	LifeSpan    int    `json:"life_span"`
}
