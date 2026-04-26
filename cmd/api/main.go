package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/casimiroarruda/go-tide-table-api/internal/platform/http/handlers"
	"github.com/casimiroarruda/go-tide-table-api/internal/platform/storage/postgresql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Error while opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	log.Println("✅ Successfully connected to Database")

	schema := os.Getenv("DATABASE_SCHEMA")
	if schema == "" {
		log.Fatal("DATABASE_SCHEMA is not set")
	}

	safeSchema := quoteIdentifier(schema)
	_, err = db.Exec(fmt.Sprintf("SET search_path TO %s, public", safeSchema))
	if err != nil {
		log.Fatalf("❌ Erro ao definir o schema: %v", err)
	}
	log.Println("📍 Schema configurado com sucesso!")

	locationRepo := postgresql.NewLocationRepo(db)
	locationHandler := handlers.NewLocationHandler(locationRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API está online 🌊"))
	})
	r.Route("/api", func(r chi.Router) {
		r.Get("/locations", locationHandler.GetLocations)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor rodando na porta %s", port)
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("Erro ao listar rotas: %s\n", err.Error())
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func quoteIdentifier(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
