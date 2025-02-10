package routes

import (
	"mock-api-server/handlers"
	"mock-api-server/middlewares"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes the router
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.LoggingMiddleware)

	r.HandleFunc("/upload", handlers.UploadMockEndpoint).Methods("POST")
	r.HandleFunc("/list", handlers.ListMockEndpoints).Methods("GET")
	r.HandleFunc("/delete/{id}", handlers.DeleteMockEndpoint).Methods("DELETE")
	r.PathPrefix("/").HandlerFunc(handlers.ServeMockEndpoint) // Catch-all for dynamic mocks

	return r
}
