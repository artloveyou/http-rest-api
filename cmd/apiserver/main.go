package main

import (
	"github.com/artloveyou/http-rest-api/internal/app/apiserver"
	"log"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
