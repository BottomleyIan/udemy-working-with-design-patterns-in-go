package models

import (
	"context"
	"time"
)

func (m *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
	  	SELECT 
  		id, 
  		breed, 
  		weight_low_lbs, 
  		weight_high_lbs,
  		cast(((weight_low_lbs + weight_high_lbs)/2)as unsigned) as average_weight,
  		lifespan,
  		coalesce(details, ''),
  		coalesce(alternate_names, ''),
  		coalesce(geographic_origin, '')
  	FROM dog_breeds
  	ORDER by breed
	`
	var breeds []*DogBreed
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b DogBreed
		err := rows.Scan(
			&b.ID,
			&b.Breed,
			&b.WeightLowLbs,
			&b.WeightHightLbs,
			&b.AverageWeight,
			&b.LifeSpan,
			&b.Details,
			&b.AlternativeNames,
			&b.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &b)
	}
	return breeds, nil
}
