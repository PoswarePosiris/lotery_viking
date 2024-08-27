package server

import (
	"encoding/json"
	"log"
	"lotery_viking/internal/server/middleware"
	"net/http"
)

const jsonContentType = "application/json"

func (s *Server) RegisterRoutes() http.Handler {

	router := http.NewServeMux()
	// public route
	router.HandleFunc("/", s.HelloWorldHandler)

	// protected route
	s.addProtectedRoute(router, "/test", s.HelloWorldHandler)

	// Health check
	router.HandleFunc("/health", s.healthHandler)

	return router
}

func (s *Server) addProtectedRoute(router *http.ServeMux, path string, handler http.HandlerFunc) {
	protectedHandler := middleware.CheckAPIKey(handler)
	router.Handle(path, protectedHandler)
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonContentType)
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
