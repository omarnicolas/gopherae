package listing

import (
	"errors"
)

// ErrNotFound is used when a gopher could not be found.
var ErrNotFound = errors.New("gopher not found")

// Repository provides access to the gopher and review storage.
type Repository interface {
	// GetGopher returns the gopher with given ID.
	GetGopher(int) (Gopher, error)
	// GetAllGophers returns all gophers saved in storage.
	GetAllGophers() []Gopher
	// GetAllReviews returns a list of all reviews for a given gopher ID.
	GetAllReviews(int) []Review
}

// Service provides gopher and review listing operations.
type Service interface {
	GetGopher(int) (Gopher, error)
	GetGophers() []Gopher
	GetGopherReviews(int) []Review
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetGophers returns all gophers
func (s *service) GetGophers() []Gopher {
	return s.r.GetAllGophers()
}

// GetGopher returns a gopher
func (s *service) GetGopher(id int) (Gopher, error) {
	return s.r.GetGopher(id)
}

// GetGopherReviews returns all requests for a gopher
func (s *service) GetGopherReviews(gopherID int) []Review {
	return s.r.GetAllReviews(gopherID)
}
