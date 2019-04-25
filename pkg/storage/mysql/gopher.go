package mysql

import "time"

// Gopher defines the storage form of a gopher
type Gopher struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}
