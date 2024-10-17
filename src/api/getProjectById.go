package api

import (
	"encoding/json"
	"infrastructure-catalog-backend/src/models"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProjectById(w http.ResponseWriter, r *http.Request) {

	dbname := os.Getenv("DB_NAME")

	// Start a tracing span
	ctx, span := tracer.Start(r.Context(), "GetProjectById")
	defer span.End()

	// Retrieve MongoDB Client from context
	client := ctx.Value("mongoClient").(*mongo.Client)
	collection := client.Database(dbname).Collection("projects")

	// Get the project ID from the query parameters
	projectID := r.URL.Query().Get("id")
	if projectID == "" {
		http.Error(w, "Missing project ID", http.StatusBadRequest)
		log.Println("Project ID is required")
		return
	}

	// Query the MongoDB collection for the project by ID
	var rawProject map[string]interface{}
	err := collection.FindOne(ctx, bson.M{"id": projectID}).Decode(&rawProject)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Project not found", http.StatusNotFound)
			log.Println("Project not found with ID:", projectID)
		} else {
			http.Error(w, "Failed to retrieve project", http.StatusInternalServerError)
			log.Println("Error retrieving project from database:", err)
		}
		return
	}

	// Parse the raw data into the Project struct
	var project models.Project
	project.ID = rawProject["id"].(string)
	project.Name = rawProject["name"].(string)
	project.Description = rawProject["description"].(string)
	project.JSONData = rawProject["jsondata"].(map[string]interface{})

	// Return the project as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		log.Println("Error encoding project to JSON:", err)
		http.Error(w, "Failed to encode project data", http.StatusInternalServerError)
		return
	}
	log.Println("Project retrieved successfully with ID:", projectID)
}
