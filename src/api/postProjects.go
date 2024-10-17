package api

import (
	"encoding/json"
	"fmt"
	"infrastructure-catalog-backend/src/models"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {

	dbname := os.Getenv("DB_NAME")
	// Start a tracing span (assuming tracer is defined)
	ctx, span := tracer.Start(r.Context(), "CreateProject")
	defer span.End()

	// Retrieve MongoDB Client from context
	client := ctx.Value("mongoClient").(*mongo.Client)
	collection := client.Database(dbname).Collection("projects")

	if r.Method == http.MethodPost {
		var project models.Project

		// Decode the incoming JSON
		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			log.Println("Invalid JSON: ", err)
			return
		}

		// Generate a new UUID for the project's ID
		project.ID = uuid.New().String()

		// Log the project details
		log.Println("Project ID:", project.ID)
		log.Println("Project Name:", project.Name)
		log.Println("Project Description:", project.Description)
		log.Println("Project JSONData:", project.JSONData)

		// Insert the data into the MongoDB collection
		_, err = collection.InsertOne(ctx, project)
		if err != nil {
			http.Error(w, "Failed to save project to database", http.StatusInternalServerError)
			log.Println("Failed to save project to the database: ", err)
			return
		}

		// Send a success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Project created successfully", "id": "%s"}`, project.ID)
		log.Println("Project created successfully with ID: ", project.ID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
