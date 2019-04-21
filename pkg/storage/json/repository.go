package json

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"strconv"
	"time"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/omarnicolas/gopherae/pkg/adding"
	"github.com/omarnicolas/gopherae/pkg/listing"
	"github.com/omarnicolas/gopherae/pkg/reviewing"
)

const (
	// dir defines the name of the directory where the files are stored
	dir = "/data/"

	// CollectionGopher identifier for the JSON collection of gophers
	CollectionGopher = "gophers"
	// CollectionReview identifier for the JSON collection of reviews
	CollectionReview = "reviews"
)

// Storage stores gopher data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// AddGopher saves the given gopher to the repository
func (s *Storage) AddGopher(b adding.Gopher) error {

	existingGophers := s.GetAllGophers()
	for _, e := range existingGophers {
		if b.Name == e.Name {
			return adding.ErrDuplicate
		}
	}

	newB := Gopher{
		ID:        len(existingGophers) + 1,
		Created:   time.Now(),
		Name:      b.Name,
		ShortDesc: b.ShortDesc,
	}

	resource := strconv.Itoa(newB.ID)
	if err := s.db.Write(CollectionGopher, resource, newB); err != nil {
		return err
	}
	return nil
}

// AddReview saves the given review in the repository
func (s *Storage) AddReview(r reviewing.Review) error {

	var gopher Gopher
	if err := s.db.Read(CollectionGopher, strconv.Itoa(r.GopherID), &gopher); err != nil {
		return listing.ErrNotFound
	}

	created := time.Now()
	newR := Review{
		ID:        fmt.Sprintf("%d_%s_%s_%d", r.GopherID, r.FirstName, r.LastName, created.Unix()),
		Created:   created,
		GopherID:  r.GopherID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
	}

	if err := s.db.Write(CollectionReview, newR.ID, r); err != nil {
		return err
	}

	return nil
}

// Get returns a gopher with the specified ID
func (s *Storage) GetGopher(id int) (listing.Gopher, error) {
	var g Gopher
	var gopher listing.Gopher

	var resource = strconv.Itoa(id)

	if err := s.db.Read(CollectionGopher, resource, &g); err != nil {
		// err handling omitted for simplicity
		return gopher, listing.ErrNotFound
	}

	gopher.ID = g.ID
	gopher.Name = g.Name
	gopher.ShortDesc = g.ShortDesc
	gopher.Created = g.Created

	return gopher, nil
}

// GetAll returns all gophers
func (s *Storage) GetAllGophers() []listing.Gopher {
	list := []listing.Gopher{}

	records, err := s.db.ReadAll(CollectionGopher)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, r := range records {
		var g Gopher
		var gopher listing.Gopher

		if err := json.Unmarshal([]byte(r), &g); err != nil {
			// err handling omitted for simplicity
			return list
		}

		gopher.ID = g.ID
		gopher.Name = g.Name
		gopher.ShortDesc = g.ShortDesc
		gopher.Created = g.Created

		list = append(list, gopher)
	}

	return list
}

// GetAll returns all reviews for a given gopher
func (s *Storage) GetAllReviews(gopherID int) []listing.Review {
	list := []listing.Review{}

	records, err := s.db.ReadAll(CollectionReview)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, b := range records {
		var r Review

		if err := json.Unmarshal([]byte(b), &r); err != nil {
			// err handling omitted for simplicity
			return list
		}

		if r.GopherID == gopherID {
			var review listing.Review

			review.ID = r.ID
			review.GopherID = r.GopherID
			review.FirstName = r.FirstName
			review.LastName = r.LastName
			review.Score = r.Score
			review.Text = r.Text
			review.Created = r.Created

			list = append(list, review)
		}
	}

	return list
}
