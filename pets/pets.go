package pets

type Pet struct {
	Species          string `json:"species"`
	Breed            string `json:"breed"`
	Description      string `json:"description,omitempty"`
	GeographicOrigin string `json:"geographic_origin,omitempty"`
	Color            string `json:"color,omitempty"`
	MinWeight        int    `json:"min_weight,omitempty"`
	MaxWeight        int    `json:"max_weight,omitempty"`
	AverageWeight    int    `json:"average_weight,omitempty"`
	Weight           int    `json:"weight,omitempty"`
	LifeSpan         int    `json:"lifespan,omitempty"`
	Age              int    `json:"age,omitempty"`
	AgeEstimated     bool   `json:"age_estimated,omitempty"`
}
