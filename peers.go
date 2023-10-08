package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var peers []string = []string{
	"http://localhost:9090",
	"http://localhost:9091",
	"http://localhost:9092",
	"http://localhost:9093",
}

func getCurrentPeer() (string, error) {
	for _, peer := range peers {
		if strings.Contains(peer, port) {
			return peer, nil
		}
	}
	return "", errors.New("Peer not found in the list")
}

func InformPeersAboutFavoriteNumber(favoriteNumber int) error {
	currentPeer, err := getCurrentPeer()
	if err != nil {
		return err
	}

	var postingFavoriteNumberError error = nil
	for _, peer := range peers {
		if peer != currentPeer {
			err = postFavoriteNumber(peer, favoriteNumber)
			if err != nil {
				log.Println(err)
				postingFavoriteNumberError = err
			}
		}
	}

	if postingFavoriteNumberError != nil {
		return postingFavoriteNumberError
	}

	return nil
}

func postFavoriteNumber(peer string, favoriteNumber int) error {
	postUrl := peer + "/favorite_number"

	postData := []byte(`{
		"number": ` + strconv.Itoa(favoriteNumber) + `,
		"from_peer": true
	}`)

	r, err := http.NewRequest("POST", postUrl, strings.NewReader(string(postData)))
	if err != nil {
		return err
	}

	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Couldn't post favorite number to " + peer + " with status code " + strconv.Itoa(response.StatusCode))
	}

	return nil
}
