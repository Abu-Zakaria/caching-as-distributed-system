package main

import (
	"log"
	"net/http"
	"os"
)

var port string = "9090"

func main() {
	arguments := os.Args[1:]

	if len(arguments) > 0 {
		port = arguments[0]
	}

	http.HandleFunc("/favorite_number", HandleFavoriteNumberSaving)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
