package main

import (
	"log"

	"github.com/firdasafridi/example-dependency-injection/country"
)

func main() {
	handler := newHandler(&country.HTTPCountry{})
	handler.printAllCountry()
}

// Handler the repository
type Handler struct {
	Country country.Repository
}

func newHandler(repoCountry country.Repository) *Handler {
	return &Handler{
		Country: repoCountry,
	}
}

func (handler *Handler) printAllCountry() {
	listCountry, err := handler.Country.GetAll()
	if err != nil {
		log.Println("[main] Failed get all country:", err)
		return
	}

	for _, country := range listCountry {
		log.Println(country.Name)
	}
}
