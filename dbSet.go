package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, err = gorm.Open(sqlite.Open("tweets.db"), &gorm.Config{})
