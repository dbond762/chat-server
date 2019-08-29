package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Setup(userHandler *UserHandler, port int) {
	r := chi.NewRouter()

	CORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(CORS.Handler)

	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/users/add", userHandler.Add)

	log.Printf("Server run on http://localhost:%d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		log.Fatal("HTTP: err on ListenAndServe: ", err)
	}
}
