package routes

import (
	"portfolio-backend/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handlers *handlers.Handlers) {
	// Health check
	router.GET("/api/health", handlers.HealthCheck)

	// API endpoints
	router.GET("/api/profile", handlers.GetProfile)
	router.GET("/api/projects", handlers.GetProjects)
	router.GET("/api/skills", handlers.GetSkills)
	router.GET("/api/experiences", handlers.GetExperiences)

	// About section endpoints
	router.GET("/api/about-cards", handlers.GetAboutCards)
	router.GET("/api/learning-items", handlers.GetLearningItems)
	router.GET("/api/goals", handlers.GetGoals)
	router.GET("/api/fun-facts", handlers.GetFunFacts)

	// Upload endpoint
	router.POST("/api/upload/image", handlers.UploadImage)

	// Serve static files from uploads directory
	router.Static("/uploads", "./uploads")
}
