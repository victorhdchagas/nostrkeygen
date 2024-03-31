package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Build() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/*", http.FileServer(http.Dir("website/dist")).ServeHTTP)
	apiRouter := chi.NewRouter()
	apiRouter.Get("/metrics", GetMetrics)
	apiRouter.Post("/genkeys", GetKeys)
	apiRouter.Get("/decode/nprofile/{nprofile}", DecodeNProfile)

	router.Mount("/api", apiRouter)

	return router
}
