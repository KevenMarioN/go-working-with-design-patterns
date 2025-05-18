package models

import (
	"context"
	"time"
)

func (c *CatBreed) allCatBreeds() ([]*CatBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
				cast(((weight_low_lbs + weight_high_lbs) / 2) AS unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''),
				coalesce(geographic_origin, '')
				from dog_breeds order by breed`
	var breeds []*CatBreed

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cb CatBreed
		if err = rows.Scan(
			&cb.ID,
			&cb.Breed.Breed,
			&cb.WeightLowLbs,
			&cb.WeightHighLbs,
			&cb.AverageWeight,
			&cb.Lifespan,
			&cb.AlternateNames,
			&cb.GeographOrigin,
			&cb.GeographOrigin,
		); err != nil {
			return nil, err
		}
		breeds = append(breeds, &cb)
	}

	return breeds, nil
}
