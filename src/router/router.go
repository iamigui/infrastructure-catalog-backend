package router

import (
	"infrastructure-catalog-backend/src/api"
	"infrastructure-catalog-backend/src/middleware"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	err := godotenv.Load("../.env")

	if err != nil {
		log.Println(err)
		log.Fatalf("Error loading .env file")
	}

	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")

	r.Use(middleware.ConnectToDatabase(dbname, dbuser, dbpass, dbhost, dbport))

	r.HandleFunc("/GetInfra", api.GetInfraBase).Methods(("GET"))

	log.Println("Server running on localhost:8000")

	return r
}
