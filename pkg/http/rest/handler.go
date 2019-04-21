package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/omarnicolas/gopherae/pkg/adding"
	"github.com/omarnicolas/gopherae/pkg/listing"
	"github.com/omarnicolas/gopherae/pkg/reviewing"
)

func Handler(a adding.Service, l listing.Service, r reviewing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/gophers", getGophers(l))
	router.GET("/gophers/:id", getGopher(l))
	router.GET("/gophers/:id/reviews", getGopherReviews(l))

	router.POST("/gophers", addGopher(a))
	router.POST("/gophers/:id/reviews", addGopherReview(r))

	return router
}

// addGopher returns a handler for POST /gophers requests
func addGopher(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newGopher adding.Gopher
		err := decoder.Decode(&newGopher)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.AddGopher(newGopher)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New gopher added.")
	}
}

// addGopherReview returns a handler for POST /gophers/:id/reviews requests
func addGopherReview(s reviewing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid Gopher ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		var newReview reviewing.Review
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&newReview); err != nil {
			http.Error(w, "Failed to parse review", http.StatusBadRequest)
		}

		newReview.GopherID = ID

		s.AddGopherReview(newReview)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New gopher review added.")
	}
}

// getGophers returns a handler for GET /gophers requests
func getGophers(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetGophers()
		json.NewEncoder(w).Encode(list)
	}
}

// getGopher returns a handler for GET /gophers/:id requests
func getGopher(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid gopher ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		gopher, err := s.GetGopher(ID)
		if err == listing.ErrNotFound {
			http.Error(w, "The gopher you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gopher)
	}
}

// getGopherReviews returns a handler for GET /gophers/:id/reviews requests
func getGopherReviews(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid gopher ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		reviews := s.GetGopherReviews(ID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviews)
	}
}
