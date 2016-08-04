package main

import (
	"github.com/shimastripe/go-api-sokushukai/db"
	"github.com/shimastripe/go-api-sokushukai/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	s.Run(":8080")
}
