package api

import (
	"encoding/json"
	"fmt"
	"infrastructure-catalog-backend/src/models"
	"log"
	"net/http"
	"os"
	"regexp"

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

		// Validate the project name and description
		if !validateInput(project.Name) || !validateInput(project.Description) {
			log.Println("Invalid input: Name and description must contain valid characters.")
			http.Error(w, "Invalid input: Name and description must contain valid characters.", http.StatusBadRequest)
			return
		}

		// Additional validation on jsonData if necessary
		if project.JSONData == nil {
			log.Println("Invalid JSON data: jsonData cannot be empty.")
			http.Error(w, "Invalid JSON data: jsonData cannot be empty.", http.StatusBadRequest)
			return
		}

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

// validateInput checks if a string contains only valid characters
func validateInput(input string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9\s.,-]*$`)
	return regex.MatchString(input) && len(input) > 0
}
