package main

import "github.com/omarnicolas/gopherae/pkg/reviewing"

var DefaultReviews = []reviewing.Review{
	{GopherID: 1, FirstName: "John", LastName: "Buffalo", Score: 5, Text: "The best super Gopher."},
	{GopherID: 2, FirstName: "Chris", LastName: "Martini", Score: 1, Text: "I would SO NOT see this ever again."},
	{GopherID: 1, FirstName: "Stephen", LastName: "King", Score: 4, Text: "Was pretty good!"},
	{GopherID: 2, FirstName: "Randy", LastName: "Green", Score: 2, Text: "Wasn't that great."},
	{GopherID: 1, FirstName: "Maria", LastName: "Smith", Score: 5, Text: "AMAZING!"},
	{GopherID: 2, FirstName: "Veronica", LastName: "Green", Score: 5, Text: "Super!"},
}
