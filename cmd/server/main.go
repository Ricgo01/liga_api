package main

import (
	"log"
	"net/http"

	"liga-api/internal/handlers"
	"liga-api/internal/storage"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	storage.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/api/matches", handlers.GetMatches).Methods("GET")
	r.HandleFunc("/api/matches/{id}", handlers.GetMatchByID).Methods("GET")
	r.HandleFunc("/api/matches", handlers.CreateMatch).Methods("POST")
	r.HandleFunc("/api/matches/{id}", handlers.UpdateMatch).Methods("PUT")
	r.HandleFunc("/api/matches/{id}", handlers.DeleteMatch).Methods("DELETE")
	r.HandleFunc("/api/matches/{id}/goals", handlers.UpdateGoals).Methods("PATCH")
	r.HandleFunc("/api/matches/{id}/yellowcards", handlers.UpdateYellowCards).Methods("PATCH")
	r.HandleFunc("/api/matches/{id}/redcards", handlers.UpdateRedCards).Methods("PATCH")
	r.HandleFunc("/api/matches/{id}/extratime", handlers.UpdateExtraTime).Methods("PATCH")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(r)

	log.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
