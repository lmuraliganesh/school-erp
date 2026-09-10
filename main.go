package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"school-erp/database"
	"school-erp/handlers"

	appmw "school-erp/middleware"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Connect to PostgreSQL and initialize tables
	database.ConnectDB()
	database.InitUserTable()
	database.InitClassTable()
	database.InitSubjectTable()
	database.InitStudentProfileTable()
	database.InitTimetableTable()

	// Initialize Chi router
	r := chi.NewRouter()
	// Allow requests from our React frontend
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:5174"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.With(appmw.AuthMiddleware, appmw.AdminOnlyMiddleware).Post("/api/users", handlers.CreateUserHandler)
	r.With(appmw.AuthMiddleware, appmw.AdminOnlyMiddleware).Get("/api/users", handlers.GetUsersHandler)
	r.With(appmw.AuthMiddleware, appmw.AdminOnlyMiddleware).Post("/api/classes", handlers.CreateClassHandler)
	r.With(appmw.AuthMiddleware, appmw.AdminOnlyMiddleware).Post("/api/subjects", handlers.CreateSubjectHandler)
	r.With(appmw.AuthMiddleware, appmw.AdminOnlyMiddleware).Post("/api/students", handlers.EnrollStudentHandler)
	r.With(appmw.AuthMiddleware, appmw.AdminOnlyMiddleware).Get("/api/students", handlers.GetStudentsHandler)
	r.With(appmw.AuthMiddleware, appmw.AdminOnlyMiddleware).Post("/api/timetables", handlers.CreateTimetableHandler)
	r.With(appmw.AuthMiddleware, appmw.AdminOnlyMiddleware).Get("/api/timetables", handlers.GetTimetableHandler)
	r.Post("/api/login", handlers.LoginHandler)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "School ERP API is up and running with PostgreSQL!",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is starting on http://localhost:%s...\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
