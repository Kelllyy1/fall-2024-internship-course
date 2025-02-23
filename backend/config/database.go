/* In this file, I attempt to connect to my database.
 */

/* I want to be able to connect to these databases:

courses
goals

*/

package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to the database
func ConnectDB() *mongo.Client {

	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set the connection string, imported from the .env file
	connectionString := os.Getenv("MONGO_URI")
	if connectionString == "" {
		log.Fatalf("Connection string is empty")

	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString).SetAppName("final-project"))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping the database to ensure connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	fmt.Println("Connected to MongoDB!")

	return client
}

// Set the connection string, imported from the .env file
// Attempt to connect to MongoDB
// Ping the database to ensure connection
// Log success or failure message
// Return the connection
