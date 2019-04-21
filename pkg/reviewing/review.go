package reviewing

// Review defines a gopher review
type Review struct {
	GopherID  int    `json:"gopher_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Score     int    `json:"score"`
	Text      string `json:"text"`
}
