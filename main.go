// @title Portfolio Backend API
// @version 1.0
// @description API for portfolio website backend
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https

package main

import (
	"log"
	"net/http"
	"os"
	"portfolio-backend/database"
	"portfolio-backend/handlers"
	"portfolio-backend/repository"
	"portfolio-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Import swagger docs
	_ "portfolio-backend/docs"
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
	frontendURL := os.Getenv("FRONTEND_URL")
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3001", frontendURL}
	config.AllowCredentials = true
	log.Println("FRONTEND_URL", frontendURL)
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Register routes
	routes.RegisterRoutes(r, h)

	// Swagger route
	r.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redirect /api-docs to /api-docs/index.html
	r.GET("/api-docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api-docs/index.html")
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("Swagger UI available at http://localhost:%s/api-docs", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
