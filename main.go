package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"GoWebApiSQL/tweet"
	"net/http"
)

var db, err = gorm.Open(sqlite.Open("tweets.db"), &gorm.Config{})

var tweets []Twit

type Twit tweet.Tweet

func Migrate() {
	db.AutoMigrate(&Twit{})
}

func connectDB() *gorm.DB {
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

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

func getAllTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := []map[string]interface{}{}
	db.Table("tweets").Find(&result)
	json.NewEncoder(w).Encode(&result)

}

func getTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tweet Twit
	tweet.TweetID, _ = uuid.FromString(params["tweet_id"])
	result := map[string]interface{}{}
	db.Table("tweets").Where("tweet_id = ?", tweet.TweetID).Find(&result)
	json.NewEncoder(w).Encode(&result)
}

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
