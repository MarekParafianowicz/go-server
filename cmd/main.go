package main

import (
	"log"

	"github.com/marekparafianowicz/go-server/pkg/repository"
	"github.com/marekparafianowicz/go-server/pkg/server"
)

func main() {
	repository, err := repository.New("postgres", "postgres://postgres:password@postgres/go-server?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(repository)
	s.Run()
}
