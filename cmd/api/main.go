package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"twitter-clone-api/internal/database"
	"twitter-clone-api/internal/handlers"
	"twitter-clone-api/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type App struct {
    router *chi.Mux
    db     *sql.DB
}

func NewApp() *App {
    // Chargement des variables d'environnement
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found")
    }

    db, err := database.NewPostgresDB(database.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
    })
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    router := chi.NewRouter()
    
    // Middlewares de base
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)
    router.Use(middleware.RequestID)
    router.Use(middleware.RealIP)

    return &App{
        router: router,
        db:     db,
    }
}

func (app *App) SetupRoutes() {
    // Création du repository utilisateur
    userRepo := database.NewUserRepository(app.db)
    // Initialisation des handlers avec le repository
    userHandler := handlers.NewUserHandler(services.NewUserService(userRepo))

    // Routes
    app.router.Route("/api/v1", func(r chi.Router) {
        // Users routes
        r.Route("/users", func(r chi.Router) {
            r.Post("/", userHandler.Create)
            r.Get("/{userID}", userHandler.GetOne)
            r.Put("/{userID}", userHandler.Update)
        })
    })
}

func main() {
    app := NewApp()
    app.SetupRoutes()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on :%s", port)
    if err := http.ListenAndServe(":"+port, app.router); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}