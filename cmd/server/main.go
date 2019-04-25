package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/omarnicolas/gopherae/pkg/adding"
	"github.com/omarnicolas/gopherae/pkg/http/rest"
	"github.com/omarnicolas/gopherae/pkg/listing"
	"github.com/omarnicolas/gopherae/pkg/reviewing"
	"github.com/omarnicolas/gopherae/pkg/storage/json"
	"github.com/omarnicolas/gopherae/pkg/storage/memory"
	"github.com/omarnicolas/gopherae/pkg/storage/mysql"
)

// StorageType defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
	// MySQL will store data in DB
	MySQL
)

func main() {

	// set up storage
	storageType := MySQL // this could be a flag; hardcoded here for simplicity

	var adder adding.Service
	var lister listing.Service
	var reviewer reviewing.Service

	switch storageType {
	case Memory:
		s := new(memory.Storage)

		adder = adding.NewService(s)
		lister = listing.NewService(s)
		reviewer = reviewing.NewService(s)

	case JSON:
		// error handling omitted for simplicity
		s, _ := json.NewStorage()

		adder = adding.NewService(s)
		lister = listing.NewService(s)
		reviewer = reviewing.NewService(s)

	case MySQL:
		// error handling omitted for simplicity
		s, _ := mysql.NewConnection()

		adder = adding.NewService(s)
		lister = listing.NewService(s)
		reviewer = reviewing.NewService(s)
	}

	// set up the HTTP server
	router := rest.Handler(adder, lister, reviewer)

	fmt.Println("The gopher server is on now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
