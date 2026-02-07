package main

import (
	"go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.Ldate|log.Ltime)

	server := server.NewServer(logger)
	err := server.Server.ListenAndServe()
	if err != nil{
		logger.Fatal(err)
	}
}
