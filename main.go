package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var tweet []Tweet

var db, err = gorm.Open(sqlite.Open("tweets.db"), &gorm.Config{})

type Tweet struct {
	TweetID           uuid.UUID `gorm:"type:uuid;primary_key;not null; json:"Tweetid"`
	Username          string    `sql:"type:json" json:"Username"`
	ContentText       string    `sql:"type:json" json:"Contenttext"`
	ContentAttachment string    `sql:"type:json" json:"Contentattachment"`
}

func connectDB() *gorm.DB {
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func getTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := []map[string]interface{}{}
	db.Table("tweets").Find(&result)
	json.NewEncoder(w).Encode(&result)
	return
}

func main() {
	r := mux.NewRouter()
	connectDB()

	r.HandleFunc("/getTweet", getTweet).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
