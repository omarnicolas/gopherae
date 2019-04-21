package main

import (
	"fmt"

	"github.com/omarnicolas/gopherae/pkg/adding"
	"github.com/omarnicolas/gopherae/pkg/reviewing"
	"github.com/omarnicolas/gopherae/pkg/storage/json"
)

func main() {

	var adder adding.Service
	var reviewer reviewing.Service

	// error handling omitted for simplicity
	s, _ := json.NewStorage()

	adder = adding.NewService(s)
	reviewer = reviewing.NewService(s)

	// add some sample data
	adder.AddSampleGophers(DefaultGophers)
	reviewer.AddSampleReviews(DefaultReviews)

	fmt.Println("Finished adding sample data.")
}
