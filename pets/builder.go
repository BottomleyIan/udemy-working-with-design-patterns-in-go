package pets

import "errors"

type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(s string) *Pet
	SetDescription(s string) *Pet
	SetGeographicOrigin(s string) *Pet
	SetColor(s string) *Pet
	SetMinWeight(i int) *Pet
	SetMaxWeight(i int) *Pet
	SetWeight(i int) *Pet
	SetLifespan(i int) *Pet
	SetAge(i int) *Pet
	SetAgeEstimated(b bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() PetInterface {
	return &Pet{}
}

func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

func (p *Pet) SetBreed(s string) *Pet {
	p.Breed = s
	return p
}

func (p *Pet) SetDescription(s string) *Pet {
	p.Description = s
	return p
}

func (p *Pet) SetGeographicOrigin(s string) *Pet {
	p.GeographicOrigin = s
	return p
}

func (p *Pet) SetColor(s string) *Pet {
	p.Color = s
	return p
}

func (p *Pet) SetMinWeight(s int) *Pet {
	p.MinWeight = s
	return p
}

func (p *Pet) SetMaxWeight(s int) *Pet {
	p.MaxWeight = s
	return p
}

func (p *Pet) SetWeight(s int) *Pet {
	p.Weight = s
	return p
}

func (p *Pet) SetLifespan(s int) *Pet {
	p.LifeSpan = s
	return p
}

func (p *Pet) SetAge(s int) *Pet {
	p.Age = s
	return p
}

func (p *Pet) SetAgeEstimated(b bool) *Pet {
	p.AgeEstimated = b
	return p
}

func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("minimum weight must be less than maximum weight")
	}
	p.AverageWeight = (p.MinWeight + p.AverageWeight) / 2
	return p, nil
}
