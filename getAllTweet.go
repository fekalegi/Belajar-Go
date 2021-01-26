package main

import (
	"encoding/json"
	"net/http"
)

func getAllTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := []map[string]interface{}{}
	db.Table("tweets").Find(&result)
	json.NewEncoder(w).Encode(&result)

}