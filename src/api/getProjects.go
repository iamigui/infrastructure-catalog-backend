package api

import (
	"encoding/json"
	"infrastructure-catalog-backend/src/models"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo" // Import options for Find
	"go.opentelemetry.io/otel"
)

const name = "otel-collector"

var (
	tracer = otel.Tracer(name)
)

func GetProjectsBase(w http.ResponseWriter, r *http.Request) {

	dbname := os.Getenv("DB_NAME")
	// Start a tracing span
	ctx, span := tracer.Start(r.Context(), "GetProjectsBase")
	defer span.End()

	// Retrieve MongoDB Client from context
	client := ctx.Value("mongoClient").(*mongo.Client)
	collection := client.Database(dbname).Collection("projects")

	// Query the collection to retrieve all projects
	cursor, err := collection.Find(ctx, bson.D{}) // mongo.D{} is used for an empty filter
	if err != nil {
		log.Println("Error querying the MongoDB database:", err)
		http.Error(w, "Internal Server Error querying the database", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Prepare projects to send as JSON
	var projectsResult []models.Project

	for cursor.Next(ctx) {
		var project models.Project
		if err := cursor.Decode(&project); err != nil {
			log.Println("Error decoding document:", err)
			http.Error(w, "Internal Server Error decoding document", http.StatusInternalServerError)
			return
		}
		projectsResult = append(projectsResult, project)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		http.Error(w, "Internal Server Error during cursor iteration", http.StatusInternalServerError)
		return
	}

	// Log each project
	for _, project := range projectsResult {
		log.Printf("ID: %s, Project Name: %s, Description: %s, Infrastructure: %v", project.ID, project.Name, project.Description, project.JSONData)
	}

	// Set Content-Type header and send JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(projectsResult); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal Server Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
