package main

import (
	"infrastructure-catalog-backend/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
