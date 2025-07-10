package models

import (
	"time"

	"github.com/lib/pq"
)

// StringArray type for PostgreSQL text[] arrays
type StringArray pq.StringArray

// Profile model
type Profile struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Location  string    `json:"location"`
	Website   string    `json:"website"`
	Github    string    `json:"github"`
	Linkedin  string    `json:"linkedin"`
	Avatar    string    `json:"avatar"`
	About     string    `json:"about"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Project model
type Project struct {
	ID           int         `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Image        string      `json:"image"`
	Technologies StringArray `json:"technologies"`
	Github       string      `json:"github"`
	Demo         string      `json:"demo"`
	Featured     bool        `json:"featured"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

// Skill model
type Skill struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	IconURL  string `json:"icon_url"`
	Category string `json:"category"`
}

// Experience model
type Experience struct {
	ID           int         `json:"id"`
	Company      string      `json:"company"`
	Position     string      `json:"position"`
	Location     string      `json:"location"`
	StartDate    time.Time   `json:"start_date"`
	EndDate      *time.Time  `json:"end_date"`
	Current      bool        `json:"current"`
	Description  string      `json:"description"`
	Technologies StringArray `json:"technologies"`
	Logo         string      `json:"logo"`
}

// AboutCard model for "What Drives Me" section
type AboutCard struct {
	ID             int         `json:"id"`
	Icon           string      `json:"icon"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	GradientColors StringArray `json:"gradient_colors"`
	DisplayOrder   int         `json:"display_order"`
	HoverRotateY   int         `json:"hover_rotate_y"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

// LearningItem model for "Currently Learning" section
type LearningItem struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Color        string    `json:"color"`
	DisplayOrder int       `json:"display_order"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Goal model for "Goals Ahead" section
type Goal struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Color        string    `json:"color"`
	DisplayOrder int       `json:"display_order"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// FunFact model for "Fun Facts" section
type FunFact struct {
	ID           int       `json:"id"`
	Emoji        string    `json:"emoji"`
	Text         string    `json:"text"`
	RotateAngle  int       `json:"rotate_angle"`
	DisplayOrder int       `json:"display_order"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
