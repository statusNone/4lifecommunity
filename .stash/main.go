package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/statusnone/4life-community/internal/content"
	"github.com/statusnone/4life-community/internal/db"
	"github.com/statusnone/4life-community/internal/render"
	"github.com/statusnone/4life-community/internal/server"
)

func main() {
	ctx := context.Background()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://4life:4life@localhost:5432/4life?sslmode=disable"
	}
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		adminPassword = "4life-admin"
	}
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		sessionSecret = "dev-secret-change-me"
	}

	store, err := db.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer store.Close()

	if err := store.Migrate(ctx); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	if err := content.Seed(ctx, store); err != nil {
		log.Fatalf("seed: %v", err)
	}

	publicTpl, err := render.NewLoader("views", "base.html").Load()
	if err != nil {
		log.Fatalf("templates: %v", err)
	}
	adminTpl, err := render.NewLoader("views", "admin_base.html").Load()
	if err != nil {
		log.Fatalf("admin templates: %v", err)
	}

	srv := server.New(server.Config{
		Store:          store,
		PublicTpl:      publicTpl,
		AdminTpl:       adminTpl,
		AdminPassword:  adminPassword,
		SessionSecret:  []byte(sessionSecret),
	})

	addr := ":" + port
	log.Printf("4life listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, srv.Routes()); err != nil {
		log.Fatal(err)
	}
}
