package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// cached for other peers
var favoriteNumber int

type FavoriteNumber struct {
	Number   int  `json:"number"`
	FromPeer bool `json:"from_peer"`
}

func HandleFavoriteNumberSaving(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	if r.Method == http.MethodGet {
		w.Write([]byte(strconv.Itoa(favoriteNumber)))
	} else if r.Method == http.MethodPost {

		decoder := json.NewDecoder(r.Body)
		var favoriteNumberData FavoriteNumber

		err := decoder.Decode(&favoriteNumberData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid favorite number"))
			log.Println(err)
			return
		}

		favoriteNumber = favoriteNumberData.Number

		w.Write([]byte("Favorite number saved"))

		if favoriteNumberData.FromPeer != true {
			log.Println("Informing peers about favorite number")
			err = InformPeersAboutFavoriteNumber(favoriteNumber)
			if err != nil {
				log.Println("Something went wrong while informing peers about favorite number")
			}
		}
	}
}
