package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectDB initializes the MongoDB connection
func ConnectDB() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("❌ MONGO_URI is not set in environment variables")
	}

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Set a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("🚨 Failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("🚨 MongoDB ping failed: %v", err)
	}

	MongoClient = client
	fmt.Println("✅ Connected to MongoDB")
}

// DisconnectDB ensures MongoDB connection is closed properly
func DisconnectDB() {
	if MongoClient != nil {
		if err := MongoClient.Disconnect(context.TODO()); err != nil {
			log.Printf("⚠️  Error disconnecting MongoDB: %v", err)
		} else {
			fmt.Println("🔌 MongoDB connection closed")
		}
	}
}

// GetCollection returns a MongoDB collection
func GetCollection(name string) *mongo.Collection {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("❌ DB_NAME is not set in environment variables")
	}
	return MongoClient.Database(dbName).Collection(name)
}
