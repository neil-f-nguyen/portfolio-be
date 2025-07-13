package models

import (
	"time"

	"github.com/lib/pq"
)

// StringArray type for PostgreSQL text[] arrays
// swagger:type array,string
type StringArray pq.StringArray

// Profile model
type Profile struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"Nguyen Anh Tu"`
	Title     string    `json:"title" example:"Frontend Developer"`
	Email     string    `json:"email" example:"tu.nguyen@example.com"`
	Phone     string    `json:"phone" example:"+84 123 456 789"`
	Location  string    `json:"location" example:"Hanoi, Vietnam"`
	Website   string    `json:"website" example:"https://tuportfolio.com"`
	Github    string    `json:"github" example:"https://github.com/tunguyen"`
	Linkedin  string    `json:"linkedin" example:"https://linkedin.com/in/tunguyen"`
	Avatar    string    `json:"avatar" example:"https://example.com/avatar.jpg"`
	About     string    `json:"about" example:"Frontend Developer with 3 years of experience"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// Project model
type Project struct {
	ID           int         `json:"id" example:"1"`
	Title        string      `json:"title" example:"E-Commerce Platform"`
	Description  string      `json:"description" example:"Modern e-commerce platform with online payment"`
	Image        StringArray `json:"image" swaggertype:"array,string" example:"https://example.com/image1.jpg,https://example.com/image2.jpg"`
	Technologies StringArray `json:"technologies" swaggertype:"array,string" example:"React,TypeScript,Node.js"`
	Github       string      `json:"github" example:"https://github.com/tunguyen/ecommerce"`
	Demo         string      `json:"demo" example:"https://ecommerce.tuportfolio.com"`
	Featured     bool        `json:"featured" example:"true"`
	CreatedAt    time.Time   `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt    time.Time   `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// Skill model
type Skill struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"React"`
	Level    string `json:"level" example:"Advanced"`
	IconURL  string `json:"icon_url" example:"https://example.com/react-icon.svg"`
	Category string `json:"category" example:"Frontend"`
}

// Experience model
type Experience struct {
	ID           int         `json:"id" example:"1"`
	Company      string      `json:"company" example:"TechCorp Solutions"`
	Position     string      `json:"position" example:"Frontend Developer"`
	Location     string      `json:"location" example:"Hanoi, Vietnam"`
	StartDate    time.Time   `json:"start_date" example:"2022-01-01T00:00:00Z"`
	EndDate      *time.Time  `json:"end_date" example:"2023-12-31T00:00:00Z"`
	Current      bool        `json:"current" example:"true"`
	Description  string      `json:"description" example:"Developed and optimized UI for web products"`
	Technologies StringArray `json:"technologies" swaggertype:"array,string" example:"React,Next.js,TypeScript"`
	Logo         string      `json:"logo" example:"https://example.com/company-logo.png"`
}

// AboutCard model for "What Drives Me" section
type AboutCard struct {
	ID             int         `json:"id" example:"1"`
	Icon           string      `json:"icon" example:"🎨"`
	Title          string      `json:"title" example:"Creative Design"`
	Description    string      `json:"description" example:"Creating beautiful interfaces and excellent user experiences"`
	GradientColors StringArray `json:"gradient_colors" swaggertype:"array,string" example:"#3B82F6,#8B5CF6,#EC4899"`
	DisplayOrder   int         `json:"display_order" example:"1"`
	HoverRotateY   int         `json:"hover_rotate_y" example:"10"`
	CreatedAt      time.Time   `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt      time.Time   `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// LearningItem model for "Currently Learning" section
type LearningItem struct {
	ID           int       `json:"id" example:"1"`
	Title        string    `json:"title" example:"Next.js 14 & App Router"`
	Color        string    `json:"color" example:"blue"`
	DisplayOrder int       `json:"display_order" example:"1"`
	CreatedAt    time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// Goal model for "Goals Ahead" section
type Goal struct {
	ID           int       `json:"id" example:"1"`
	Title        string    `json:"title" example:"Build scalable web apps"`
	Color        string    `json:"color" example:"green"`
	DisplayOrder int       `json:"display_order" example:"1"`
	CreatedAt    time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// FunFact model for "Fun Facts" section
type FunFact struct {
	ID           int       `json:"id" example:"1"`
	Emoji        string    `json:"emoji" example:"☕"`
	Text         string    `json:"text" example:"Coffee addict\\n3+ cups/day"`
	RotateAngle  int       `json:"rotate_angle" example:"-2"`
	DisplayOrder int       `json:"display_order" example:"1"`
	CreatedAt    time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}
