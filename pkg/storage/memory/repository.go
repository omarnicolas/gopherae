package memory

import (
	"fmt"
	"time"

	"github.com/omarnicolas/gopherae/pkg/adding"
	"github.com/omarnicolas/gopherae/pkg/listing"
	"github.com/omarnicolas/gopherae/pkg/reviewing"
)

// Memory storage keeps data in memory
type Storage struct {
	gophers []Gopher
	reviews []Review
}

// Add saves the given gopher to the repository
func (m *Storage) AddGopher(g adding.Gopher) error {
	for _, e := range m.gophers {
		if g.Name == e.Name {
			return adding.ErrDuplicate
		}
	}

	newG := Gopher{
		ID:        len(m.gophers) + 1,
		Created:   time.Now(),
		Name:      g.Name,
		ShortDesc: g.ShortDesc,
	}
	m.gophers = append(m.gophers, newG)

	return nil
}

// Add saves the given review in the repository
func (m *Storage) AddReview(r reviewing.Review) error {
	found := false
	for g := range m.gophers {
		if m.gophers[g].ID == r.GopherID {
			found = true
		}
	}

	if found {
		created := time.Now()
		id := fmt.Sprintf("%d_%s_%s_%d", r.GopherID, r.FirstName, r.LastName, created.Unix())

		newR := Review{
			ID:        id,
			Created:   created,
			GopherID:  r.GopherID,
			FirstName: r.FirstName,
			LastName:  r.LastName,
			Score:     r.Score,
			Text:      r.Text,
		}

		m.reviews = append(m.reviews, newR)
	} else {
		return listing.ErrNotFound
	}

	return nil
}

// Get returns a gopher with the specified ID
func (m *Storage) GetGopher(id int) (listing.Gopher, error) {
	var gopher listing.Gopher

	for i := range m.gophers {

		if m.gophers[i].ID == id {
			gopher.ID = m.gophers[i].ID
			gopher.Name = m.gophers[i].Name
			gopher.ShortDesc = m.gophers[i].ShortDesc
			gopher.Created = m.gophers[i].Created

			return gopher, nil
		}
	}

	return gopher, listing.ErrNotFound
}

// GetAll return all gophers
func (m *Storage) GetAllGophers() []listing.Gopher {
	var gophers []listing.Gopher

	for i := range m.gophers {

		gopher := listing.Gopher{
			ID:        m.gophers[i].ID,
			Name:      m.gophers[i].Name,
			ShortDesc: m.gophers[i].ShortDesc,
			Created:   m.gophers[i].Created,
		}

		gophers = append(gophers, gopher)
	}

	return gophers
}

// GetAll returns all reviews for a given gopher
func (m *Storage) GetAllReviews(gopherID int) []listing.Review {
	var list []listing.Review

	for i := range m.reviews {
		if m.reviews[i].GopherID == gopherID {
			r := listing.Review{
				ID:        m.reviews[i].ID,
				GopherID:  m.reviews[i].GopherID,
				FirstName: m.reviews[i].FirstName,
				LastName:  m.reviews[i].LastName,
				Score:     m.reviews[i].Score,
				Text:      m.reviews[i].Text,
				Created:   m.reviews[i].Created,
			}

			list = append(list, r)
		}
	}

	return list
}
