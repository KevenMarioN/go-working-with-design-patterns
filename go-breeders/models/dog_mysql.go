package models

import (
	"context"
	"time"
)

func (d *DogBreed) allDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
				cast(((weight_low_lbs + weight_high_lbs) / 2) AS unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''),
				coalesce(geographic_origin, '')
				from dog_breeds order by breed`
	var breeds []*DogBreed

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d DogBreed
		if err = rows.Scan(
			&d.ID,
			&d.Breed.Breed,
			&d.WeightLowLbs,
			&d.WeightHighLbs,
			&d.AverageWeight,
			&d.Lifespan,
			&d.AlternateNames,
			&d.GeographOrigin,
			&d.GeographOrigin,
		); err != nil {
			return nil, err
		}
		breeds = append(breeds, &d)
	}

	return breeds, nil
}
