package main

import (
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func delTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tweet Twit
	tweet.TweetID, _ = uuid.FromString(params["tweet_id"])
	err := db.Table("tweets").Delete("tweet_id ", tweet)
	if err.Error != nil {
		w.Write([]byte("Error"))
	} else {
		w.Write([]byte("Sukses"))
	}
}