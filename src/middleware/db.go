package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB is a middleware that connects to the MongoDB database.
func ConnectToMongoDB(dbname, dbuser, dbpass, dbhost, dbport string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Construct the MongoDB URI
			uri := "mongodb://" + dbuser + ":" + dbpass + "@" + dbhost + ":" + dbport + "/" + dbname

			// Create a new client and connect to the server
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel() // Ensure that the context is cancelled after we are done

			client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
			if err != nil {
				log.Printf("Error creating MongoDB client: %v", err)
				http.Error(w, "Could not create MongoDB client", http.StatusInternalServerError)
				return
			}

			// Check the connection
			if err := client.Ping(ctx, nil); err != nil {
				log.Printf("Error connecting to MongoDB: %v", err)
				http.Error(w, "Could not connect to MongoDB", http.StatusInternalServerError)
				return
			}
			defer client.Disconnect(ctx) // Ensure disconnection when done

			log.Println("Successfully connected to MongoDB.")

			// Add the MongoDB client to the request context
			r = r.WithContext(context.WithValue(r.Context(), "mongoClient", client))

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
