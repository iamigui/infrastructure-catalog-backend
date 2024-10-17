package router

import (
	"infrastructure-catalog-backend/src/api"
	"infrastructure-catalog-backend/src/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Add CORS support
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

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

	r.Use(middleware.ConnectToMongoDB(dbname, dbuser, dbpass, dbhost, dbport))

	r.HandleFunc("/getProjects", api.GetProjectsBase).Methods(("GET"))
	r.HandleFunc("/createProject", api.CreateProject).Methods(("POST"))

	handler := c.Handler(r)

	log.Println("Server running on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
	return r
}
