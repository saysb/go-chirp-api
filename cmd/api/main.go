package main

import (
	"database/sql"
	"log"
	"net/http"
	"twitter-clone-api/internal/database"
	"twitter-clone-api/internal/handlers"
	"twitter-clone-api/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
    router *chi.Mux
    db     *sql.DB
}

func NewApp() *App {
    db, err := database.NewPostgresDB(database.Config{
        Host:     "localhost",
        Port:     "5432",
        User:     "sebastiendamy",
        Password: "postgres",
        DBName:   "twitter-clone",
        SSLMode:  "disable",
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
            r.Get("/", userHandler.GetAll)
            r.Get("/{id}", userHandler.GetOne)
            r.Put("/{id}", userHandler.Update)
            r.Delete("/{id}", userHandler.Delete)
        })
    })
}

func main() {
    app := NewApp()
    app.SetupRoutes()

    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", app.router); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}