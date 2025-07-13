package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"portfolio-backend/repository"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	repo *repository.Repository
}

func NewHandlers(repo *repository.Repository) *Handlers {
	return &Handlers{repo: repo}
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the API is running
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/health [get]
func (h *Handlers) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get the user's profile information
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} models.Profile
// @Failure 500 {object} map[string]string
// @Router /api/profile [get]
func (h *Handlers) GetProfile(c *gin.Context) {
	profile, err := h.repo.GetProfile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// GetProjects godoc
// @Summary Get all projects
// @Description Get a list of all projects
// @Tags Projects
// @Accept json
// @Produce json
// @Success 200 {array} models.Project
// @Failure 500 {object} map[string]string
// @Router /api/projects [get]
func (h *Handlers) GetProjects(c *gin.Context) {
	projects, err := h.repo.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get projects"})
		return
	}
	c.JSON(http.StatusOK, projects)
}

// GetSkills godoc
// @Summary Get all skills
// @Description Get a list of all skills
// @Tags Skills
// @Accept json
// @Produce json
// @Success 200 {array} models.Skill
// @Failure 500 {object} map[string]string
// @Router /api/skills [get]
func (h *Handlers) GetSkills(c *gin.Context) {
	skills, err := h.repo.GetSkills()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get skills"})
		return
	}
	c.JSON(http.StatusOK, skills)
}

// GetExperiences godoc
// @Summary Get all experiences
// @Description Get a list of all work experiences
// @Tags Experiences
// @Accept json
// @Produce json
// @Success 200 {array} models.Experience
// @Failure 500 {object} map[string]string
// @Router /api/experiences [get]
func (h *Handlers) GetExperiences(c *gin.Context) {
	experiences, err := h.repo.GetExperiences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get experiences"})
		return
	}
	c.JSON(http.StatusOK, experiences)
}

// GetAboutCards godoc
// @Summary Get about cards
// @Description Get a list of about cards for "What Drives Me" section
// @Tags About
// @Accept json
// @Produce json
// @Success 200 {array} models.AboutCard
// @Failure 500 {object} map[string]string
// @Router /api/about-cards [get]
func (h *Handlers) GetAboutCards(c *gin.Context) {
	cards, err := h.repo.GetAboutCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get about cards"})
		return
	}
	c.JSON(http.StatusOK, cards)
}

// GetLearningItems godoc
// @Summary Get learning items
// @Description Get a list of currently learning items
// @Tags About
// @Accept json
// @Produce json
// @Success 200 {array} models.LearningItem
// @Failure 500 {object} map[string]string
// @Router /api/learning-items [get]
func (h *Handlers) GetLearningItems(c *gin.Context) {
	items, err := h.repo.GetLearningItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get learning items"})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetGoals godoc
// @Summary Get goals
// @Description Get a list of goals ahead
// @Tags About
// @Accept json
// @Produce json
// @Success 200 {array} models.Goal
// @Failure 500 {object} map[string]string
// @Router /api/goals [get]
func (h *Handlers) GetGoals(c *gin.Context) {
	goals, err := h.repo.GetGoals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get goals"})
		return
	}
	c.JSON(http.StatusOK, goals)
}

// GetFunFacts godoc
// @Summary Get fun facts
// @Description Get a list of fun facts
// @Tags About
// @Accept json
// @Produce json
// @Success 200 {array} models.FunFact
// @Failure 500 {object} map[string]string
// @Router /api/fun-facts [get]
func (h *Handlers) GetFunFacts(c *gin.Context) {
	facts, err := h.repo.GetFunFacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get fun facts"})
		return
	}
	c.JSON(http.StatusOK, facts)
}

// UploadImage godoc
// @Summary Upload an image
// @Description Upload an image file to the server
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/upload/image [post]
func (h *Handlers) UploadImage(c *gin.Context) {
	// Parse multipart form
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	// Get file from form
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file provided"})
		return
	}
	defer file.Close()

	// Validate file type
	if !isValidImageType(header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only JPG, JPEG, PNG, GIF, WebP are allowed"})
		return
	}

	// Validate file size (max 5MB)
	if header.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size too large. Maximum 5MB allowed"})
		return
	}

	// Create uploads directory if it doesn't exist
	uploadsDir := "uploads"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create uploads directory"})
		return
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), generateRandomString(8), ext)
	filePath := filepath.Join(uploadsDir, filename)

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer dst.Close()

	// Copy uploaded file to destination
	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Return success response with file URL
	fileURL := fmt.Sprintf("/uploads/%s", filename)
	c.JSON(http.StatusOK, gin.H{
		"message":  "Image uploaded successfully",
		"url":      fileURL,
		"filename": filename,
		"size":     header.Size,
	})
}

// Helper function to validate image file types
func isValidImageType(header *multipart.FileHeader) bool {
	// Check by file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	if !validExts[ext] {
		return false
	}

	// Check by MIME type
	validMimes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}

	// Get content type from header
	contentType := header.Header.Get("Content-Type")
	return validMimes[contentType]
}

// Helper function to generate random string
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(result)
}
