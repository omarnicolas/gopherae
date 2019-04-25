package adding

import (
	"errors"
)

// ErrDuplicate is used when a gopher already exists.
var ErrDuplicate = errors.New("gopher already exists")

// Service provides gopher adding operations.
type Service interface {
	AddGopher(...Gopher)
	AddSampleGophers([]Gopher)
}

// Repository provides access to gopher repository.
type Repository interface {
	// AddGopher saves a given gopher to the repository.
	AddGopher(Gopher) error
}

type service struct {
	bR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddGopher adds the given gopher(s) to the database
func (s *service) AddGopher(g ...Gopher) {

	// any validation can be done here

	for _, gopher := range g {
		_ = s.bR.AddGopher(gopher) // error handling omitted for simplicity
	}
}

// AddSampleGophers adds some sample gophers to the database
func (s *service) AddSampleGophers(g []Gopher) {

	// any validation can be done here

	for _, gg := range g {
		_ = s.bR.AddGopher(gg) // error handling omitted for simplicity
	}
}
