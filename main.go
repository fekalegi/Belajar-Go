package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var tweet []Tweet

type Tweet struct {
	TweetID uuid.UUID `gorm:"type:uuid;primary_key;not null; json:"Tweetid"`
	Username string `sql:"type:json" json:"Username"`
	ContentText string `sql:"type:json" json:"Contenttext"`
	ContentAttachment string `sql:"type:json" json:"Contentattachment"`
}

func main()  {

	db, err := gorm.Open(sqlite.Open("tweets.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	result := []map[string]interface{}{}
	db.Table("tweets").Find(&result)
	fmt.Print(result)
}
