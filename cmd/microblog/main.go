package main

import (
	"log"
	"os"
	"os/signal"

	"go-postgres-microblog-openshift/internal/data"
	"go-postgres-microblog-openshift/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// connection to the database.
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	// start the server.
	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown.
	serv.Close()
	data.Close()
}
