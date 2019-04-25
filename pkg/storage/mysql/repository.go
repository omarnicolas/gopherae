package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/omarnicolas/gopherae/pkg/adding"
	"github.com/omarnicolas/gopherae/pkg/listing"
	"github.com/omarnicolas/gopherae/pkg/reviewing"
)

// Storage stores gopher data
type Storage struct {
	db *sql.DB
}

// NewStorage returns a new mysql connection
func NewConnection() (*Storage, error) {
	var err error

	s := new(Storage)

	s.db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/db?parseTime=true") // environment variables; hardcoded here for simplicity
	if err != nil {
		return nil, err
	}

	return s, nil
}

// AddGopher saves the given gopher to the db
func (s *Storage) AddGopher(g adding.Gopher) error {

	existingGophers := s.GetAllGophers()
	for _, e := range existingGophers {
		if g.Name == e.Name {
			return adding.ErrDuplicate
		}
	}

	newG := Gopher{
		Name:      g.Name,
		ShortDesc: g.ShortDesc,
		Created:   time.Now(),
	}

	stmt, err := s.db.Prepare("INSERT gopher SET name=?, short_description=?, created=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newG.Name, newG.ShortDesc, newG.Created)
	if err != nil {
		return err
	}

	return nil
}

// AddReview saves the given review in the repository
func (s *Storage) AddReview(r reviewing.Review) error {
	var g Gopher

	row := s.db.QueryRow("SELECT id, name, short_description, created FROM gopher WHERE id=?", r.GopherID)
	if err := row.Scan(&g.ID, &g.Name, &g.ShortDesc, &g.Created); err != nil {
		return listing.ErrNotFound
	}

	newR := Review{
		GopherID:  r.GopherID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
		Created:   time.Now(),
	}

	stmt, err := s.db.Prepare("INSERT review SET gopher_id=?, first_name=?, last_name=?, score=?, text=?, created=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newR.GopherID, newR.FirstName, newR.LastName, newR.Score, newR.Text, newR.Created)
	if err != nil {
		return err
	}

	return nil
}

// Get returns a gopher with the specified ID
func (s *Storage) GetGopher(id int) (listing.Gopher, error) {
	var g Gopher
	var gopher listing.Gopher

	row := s.db.QueryRow("SELECT id, name, short_description, created FROM gopher WHERE id=?", id)

	if err := row.Scan(&g.ID, &g.Name, &g.ShortDesc, &g.Created); err != nil {
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

	rows, err := s.db.Query("SELECT id, name, short_description, created FROM gopher")
	if err != nil {
		// err handling omitted for simplicity
		return list
	}
	defer rows.Close()

	for rows.Next() {
		var g Gopher
		var gopher listing.Gopher

		if err := rows.Scan(&g.ID, &g.Name, &g.ShortDesc, &g.Created); err != nil {
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

	rows, err := s.db.Query("SELECT id, gopher_id, first_name, last_name, score, text, created FROM review")
	if err != nil {
		// err handling omitted for simplicity
		return list
	}
	defer rows.Close()

	for rows.Next() {
		var r Review

		if err := rows.Scan(&r.ID, &r.GopherID, &r.FirstName, &r.LastName, &r.Score, &r.Text, &r.Created); err != nil {
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
