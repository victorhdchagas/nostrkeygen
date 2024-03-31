package main

import (
	"fmt"
	"net/http"

	"github.com/victorhdchagas/nostrkeygen/internals/config"
	"github.com/victorhdchagas/nostrkeygen/internals/routes"
)

func main() {
	config, err := config.DotEnv()
	if err != nil {
		panic(err)
	}
	router := routes.Build()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}
	server.ListenAndServe()
}
