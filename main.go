package main

import (
	"log"
	"os"

	"forum/trans/server"
)

func main() {
	// Open a new log file
	file, err := os.Create("mylog.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new logger that writes to the log file
	logger := log.New(file, "", log.Ldate|log.Ltime)
	server.NewServer(logger)
}
