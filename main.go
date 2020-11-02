package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var db, err = gorm.Open(sqlite.Open("tweets.db"), &gorm.Config{})

var tweets []Tweet

type Tweet struct {
	TweetID           uuid.UUID `gorm:"type:uuid;primary_key;not null; json:"Tweetid"`
	Username          string    `sql:"type:json" json:"Username"`
	ContentText       string    `sql:"type:json" json:"Contenttext"`
	ContentAttachment string    `sql:"type:json" json:"Contentattachment"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         sql.NullTime `gorm:"index"`
}

func Migrate() {
	db.AutoMigrate(&Tweet{})
}

func connectDB() *gorm.DB {
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func addTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tweet Tweet
	_ = json.NewDecoder(r.Body).Decode(&tweet)
	tweet.TweetID = uuid.NewV4()
	err := db.Create(&tweet)
	if err != nil {
		w.Write([]byte("error"))
		fmt.Print(err)
	}
	w.Write([]byte("sukses"))
}

func getTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := []map[string]interface{}{}
	db.Table("tweets").Find(&result)
	json.NewEncoder(w).Encode(&result)

}

func main() {
	r := mux.NewRouter()
	connectDB()
	Migrate()

	r.HandleFunc("/Tweet", getTweet).Methods("GET")
	r.HandleFunc("/Tweet", addTweet).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
