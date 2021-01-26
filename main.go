package main

import (
	"GoWebApiSQL/tweet"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var tweets []Twit

type Twit tweet.Tweet

func main() {
	r := mux.NewRouter()
	connectDB()
	Migrate()

	r.HandleFunc("/Tweet", getAllTweet).Methods("GET")
	r.HandleFunc("/Tweet", addTweet).Methods("POST")
	r.HandleFunc("/Tweet/{tweet_id}", delTweet).Methods("DELETE")
	r.HandleFunc("/Tweet/{tweet_id}", getTweet).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
