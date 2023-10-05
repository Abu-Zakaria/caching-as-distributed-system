package main

import (
	"net/http"
)

func HandleFavoriteNumberSaving(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP Method not allowed", http.StatusMethodNotAllowed)
	}

}
