package main

import (
	"log"
	"os"
	"portfolio-backend/database"
	"portfolio-backend/handlers"
	"portfolio-backend/repository"
	"portfolio-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to database
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repository
	repo := repository.NewRepository(db)

	// Initialize handlers
	h := handlers.NewHandlers(repo)

	// Setup Gin router
	r := gin.Default()

	// Add custom recovery middleware
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3001", os.Getenv("FRONTEND_URL")}
	config.AllowCredentials = true
	log.Println("FRONTEND_URL", os.Getenv("FRONTEND_URL"))
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Register routes
	routes.RegisterRoutes(r, h)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
