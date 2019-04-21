package listing

import (
	"time"
)

// Gopher defines the properties of a gopher to be listed
type Gopher struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}
