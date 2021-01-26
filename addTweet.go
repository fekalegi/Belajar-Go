package main

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func addTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tweet Twit
	_ = json.NewDecoder(r.Body).Decode(&tweet)
	tweet.TweetID = uuid.NewV4()
	err := db.Create(&tweet)
	if err.Error != nil {
		w.Write([]byte("Error"))
	} else {
		w.Write([]byte("Sukses"))
	}
}