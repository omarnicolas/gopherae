package memory

import (
	"time"
)

// Gopher defines the properties of a gopher to be listed
type Gopher struct {
	ID        int
	Name      string
	ShortDesc string
	Created   time.Time
}
