package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mock-api-server/config"
	"mock-api-server/routes"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Initialize MongoDB connection
	config.ConnectDB()
	defer config.DisconnectDB() // Ensure cleanup on shutdown

	// Set up routes
	r := routes.SetupRoutes()

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Change this to specific origins if needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Get server port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3002" // Default port
	}

	// Start server
	fmt.Printf("🚀 Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, corsHandler.Handler(r)))
}
