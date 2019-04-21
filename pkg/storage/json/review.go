package json

import "time"

// Review defines the storage form of a gopher review
type Review struct {
	ID        string    `json:"id"`
	GopherID  int       `json:"gopher_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}
