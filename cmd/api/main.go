package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"api.bookworm.cc/routes"
)

func main() {
	logger := log.New(os.Stdout, "[BOOKWORM-BACKEND]\t", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	logger.Println("Kicking off...")

	routes := routes.NewRoutes(logger)
	
	server := &http.Server{
		Addr:         ":5000",
		Handler:      routes.API(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
		ErrorLog:     logger,
	}

	logger.Println("Listening on :5000")
	if err := server.ListenAndServe(); err != nil {
		logger.Println(err)
		os.Exit(1)
	}
}
