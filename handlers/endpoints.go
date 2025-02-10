package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"mock-api-server/config"
	"mock-api-server/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadMockEndpoint registers a new mock API dynamically
func UploadMockEndpoint(w http.ResponseWriter, r *http.Request) {
	// Read request body
	jsonData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Unmarshal into a slice of `Endpoint`
	var endpoints []models.Endpoint
	if err := json.Unmarshal(jsonData, &endpoints); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Insert all endpoints into MongoDB
	collection := config.GetCollection("endpoints")
	var docs []interface{}
	for _, endpoint := range endpoints {
		docs = append(docs, endpoint)
	}

	_, err = collection.InsertMany(context.TODO(), docs)
	if err != nil {
		http.Error(w, "Failed to save endpoints", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Endpoints registered successfully",
	})
}

// func UploadMockEndpoint(w http.ResponseWriter, r *http.Request) {
// 	var endpoint models.Endpoint
// 	if err := json.NewDecoder(r.Body).Decode(&endpoint); err != nil {
// 		// fmt.Println(">>>>>>>>", err)
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	endpoint.CreatedAt = time.Now()
// 	collection := config.GetCollection("endpoints")

// 	result, err := collection.InsertOne(context.TODO(), endpoint)
// 	if err != nil {
// 		http.Error(w, "Failed to save endpoint", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message": "Endpoint registered",
// 		"id":      result.InsertedID,
// 	})
// }

// ListMockEndpoints fetches all registered mock endpoints
func ListMockEndpoints(w http.ResponseWriter, r *http.Request) {
	collection := config.GetCollection("endpoints")

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(">>>>>>>>", err)
		http.Error(w, "Failed to fetch endpoints", http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	var endpoints []models.Endpoint
	for cur.Next(context.TODO()) {
		var endpoint models.Endpoint
		cur.Decode(&endpoint)
		endpoints = append(endpoints, endpoint)
	}

	json.NewEncoder(w).Encode(endpoints)
}

// DeleteMockEndpoint removes an endpoint by ID
func DeleteMockEndpoint(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	collection := config.GetCollection("endpoints")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		http.Error(w, "Failed to delete endpoint", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Endpoint deleted"})
}

// ServeMockEndpoint dynamically serves stored endpoints
func ServeMockEndpoint(w http.ResponseWriter, r *http.Request) {
	collection := config.GetCollection("endpoints")

	// Find the endpoint with the same URL and method
	var endpoint models.Endpoint
	err := collection.FindOne(context.TODO(), bson.M{
		"url":    r.URL.Path,
		"method": r.Method,
	}).Decode(&endpoint)

	if err != nil {
		http.Error(w, "Mock endpoint not found", http.StatusNotFound)
		return
	}

	// Set a default status code if it's invalid (0)
	if endpoint.StatusCode == 0 {
		endpoint.StatusCode = http.StatusOK
	}

	// Set response headers
	for key, value := range endpoint.Headers {
		w.Header().Set(key, value)
	}
	w.WriteHeader(endpoint.StatusCode)

	// Convert response body to JSON string if needed
	responseBody, err := json.Marshal(endpoint.ResponseBody)
	if err != nil {
		http.Error(w, "Error processing response body", http.StatusInternalServerError)
		return
	}

	// Write response body
	w.Write(responseBody)
}
