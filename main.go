package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/favorite_number", HandleFavoriteNumberSaving)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
