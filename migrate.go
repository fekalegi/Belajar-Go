package main

func Migrate() {
	db.AutoMigrate(&Twit{})
}
