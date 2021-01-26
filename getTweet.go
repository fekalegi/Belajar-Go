package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func getTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tweet Twit
	tweet.TweetID, _ = uuid.FromString(params["tweet_id"])
	result := map[string]interface{}{}
	db.Table("tweets").Where("tweet_id = ?", tweet.TweetID).Find(&result)
	json.NewEncoder(w).Encode(&result)
}