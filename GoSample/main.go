package main

import (
	"encoding/json"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"

	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{"status": "ok"}

	json.NewEncoder(w).Encode(status)
}

func buildRoutes(router *mux.Router, zipkinObjects ZipkinObjects) *mux.Router {

	router.HandleFunc("/", healthCheck).Methods("GET")
	router.HandleFunc("/people/{id}", callStarWars(zipkinObjects)).Methods("GET")

	return router
}

func main() {
	log.Println("App starting")

	router := mux.NewRouter()
	zipkinObjects := monitoring(router)
	router.Use(zipkinhttp.NewServerMiddleware(
		zipkinObjects.tracer,
		zipkinhttp.SpanName("request")), // name for request span
	)

	buildRoutes(router, zipkinObjects)
	log.Fatal(http.ListenAndServe(":8000", router))
}
