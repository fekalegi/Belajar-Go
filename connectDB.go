package main

import (
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
