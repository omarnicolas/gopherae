package reviewing

import (
	"errors"
)

// ErrNotFound is used when a gopher could not be found.
var ErrNotFound = errors.New("gopher not found")

// Repository provides access to the review storage.
type Repository interface {
	// AddReview saves a given review.
	AddReview(Review) error
}

// Service provides reviewing operations.
type Service interface {
	AddGopherReview(Review)
	AddSampleReviews([]Review)
}

type service struct {
	rR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddGopherReview saves a new gopher review in the database
func (s *service) AddGopherReview(r Review) {
	_ = s.rR.AddReview(r) // error handling omitted for simplicity
}

// AddSampleReviews adds some sample reviews to the database
func (s *service) AddSampleReviews(r []Review) {
	for _, rr := range r {
		_ = s.rR.AddReview(rr) // error handling omitted for simplicity
	}
}
