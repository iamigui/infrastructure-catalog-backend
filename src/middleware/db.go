package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// ConnectToDatabase is a middleware that connects to the database.
func ConnectToDatabase(dbname, dbuser, dbpassword, dbhost, dbport string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Construct the connection string
			psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpassword, dbname)

			// Open the database connection
			db, err := sql.Open("postgres", psqlInfo)
			if err != nil {
				log.Printf("Error connecting to database: %v", err)
				http.Error(w, "Could not connect to database", http.StatusInternalServerError)
				return
			}
			// Test the connection
			if err := db.Ping(); err != nil {
				log.Printf("Database unreachable: %v", err)
				http.Error(w, "Database unreachable", http.StatusInternalServerError)
				return
			}
			defer db.Close()

			log.Println("Successfully connected to the database.")

			// Add the database connection to the request context
			r = r.WithContext(context.WithValue(r.Context(), "db", db))

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
