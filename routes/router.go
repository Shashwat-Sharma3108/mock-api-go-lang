package routes

import (
	"mock-api-server/handlers"
	"mock-api-server/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes the router
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.LoggingMiddleware)

	// Handle CORS Preflight Requests
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}).Methods(http.MethodOptions)

	r.HandleFunc("/upload", handlers.UploadMockEndpoint).Methods("POST", "OPTIONS")
	r.HandleFunc("/list", handlers.ListMockEndpoints).Methods("GET")
	r.HandleFunc("/delete/{id}", handlers.DeleteMockEndpoint).Methods("DELETE")
	r.PathPrefix("/").HandlerFunc(handlers.ServeMockEndpoint) // Catch-all for dynamic mocks

	return r
}
