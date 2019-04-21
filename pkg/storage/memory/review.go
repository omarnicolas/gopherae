package memory

import (
	"time"
)

// Review defines a gopher review
type Review struct {
	ID        string
	GopherID  int
	FirstName string
	LastName  string
	Score     int
	Text      string
	Created   time.Time
}
