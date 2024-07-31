package main

import (
	"log"
	"net/http"

	"github.com/Skapar/task-management/internal/db"
	"github.com/Skapar/task-management/internal/handler"
	"github.com/Skapar/task-management/internal/repository"
	"github.com/Skapar/task-management/internal/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    db.InitDB()
    defer db.CloseDB()

    userRepo := &repository.UserRepository{DB: db.DB}
    taskRepo := &repository.TaskRepository{DB: db.DB}
    projectRepo := &repository.ProjectRepository{DB: db.DB}

    userHandler := &handler.UserHandler{Repo: userRepo}
    taskHandler := &handler.TaskHandler{Repo: taskRepo}
    projectHandler := &handler.ProjectHandler{Repo: projectRepo}

    r := mux.NewRouter()

    // Register routes
    routes.RegisterUserRoutes(r, userHandler)
    routes.RegisterTaskRoutes(r, taskHandler)
    routes.RegisterProjectRoutes(r, projectHandler)

    log.Fatal(http.ListenAndServe(":8000", r))
}
