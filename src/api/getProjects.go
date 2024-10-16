package api

import (
	"database/sql"
	"encoding/json"
	"infrastructure-catalog-backend/src/models"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
)

const name = "otel-collector"

var (
	tracer = otel.Tracer(name)
	// meter  = otel.Meter(name)
)

func GetProjectsBase(w http.ResponseWriter, r *http.Request) {

	// Start a tracing span
	ctx, span := tracer.Start(r.Context(), "GetProjectsBase")
	defer span.End()

	// Retrieve Database Client from context
	db := ctx.Value("db").(*sql.DB)
	rows, err := db.Query("SELECT id, name, description, json_data FROM projects")
	if err != nil {
		log.Println("Error querying the database:", err)
		http.Error(w, "Internal Server Error authenticating in Handler", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Prepare projects to send as JSON
	var projectsResult []models.Project
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.JSONData); err != nil {
			log.Println("Error scanning row:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		projectsResult = append(projectsResult, project)
	}

	// Log each project
	for _, project := range projectsResult {
		log.Printf("Project Name: %s, Description: %s, ID: %d, Infrastructure: %s", project.Name, project.Description, project.ID, project.JSONData)
	}

	// Set Content-Type header and send JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(projectsResult); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal Server Error Encoding JSON", http.StatusInternalServerError)
		return
	}
}
