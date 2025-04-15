package main

import (
	"fmt"

	"net/http"

	"github.com/rs/cors"
	"github.com/sohibjon7731/go-compiler/controllers"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/check", controllers.CheckSyntax)
	mux.HandleFunc("/compile", controllers.CompileHandler)
	mux.HandleFunc("/format", controllers.FormatHandler)

	handler := c.Handler(mux)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
