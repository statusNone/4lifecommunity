package main

import (
	"log"
	"net/http"
	"os"

	"github.com/statusnone/4life-community/internal/render"
	"github.com/statusnone/4life-community/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	tpl, err := render.NewLoader("views")
	if err != nil {
		log.Fatalf("templates: %v", err)
	}

	srv := server.New(tpl)

	addr := ":" + port
	log.Printf("4Life Community listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, srv.Routes()); err != nil {
		log.Fatal(err)
	}
}
