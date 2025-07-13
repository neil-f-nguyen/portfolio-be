package repository

import (
	"database/sql"
	"log"
	"portfolio-backend/database"
	"portfolio-backend/models"

	"github.com/lib/pq"
)

type Repository struct {
	db *database.DB
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{db: db}
}

// Profile methods
func (r *Repository) GetProfile() (*models.Profile, error) {
	query := `SELECT id, name, title, email, phone, location, website, github, linkedin, avatar, about, created_at, updated_at FROM profiles LIMIT 1`
	var profile models.Profile
	var phone, location, website, github, linkedin, avatar, about sql.NullString
	err := r.db.QueryRow(query).Scan(
		&profile.ID, &profile.Name, &profile.Title, &profile.Email,
		&phone, &location, &website, &github, &linkedin, &avatar, &about,
		&profile.CreatedAt, &profile.UpdatedAt,
	)
	if err != nil {
		log.Printf("[GetProfile] Scan error: %v", err)
		return nil, err
	}
	profile.Phone = phone.String
	profile.Location = location.String
	profile.Website = website.String
	profile.Github = github.String
	profile.Linkedin = linkedin.String
	profile.Avatar = avatar.String
	profile.About = about.String
	return &profile, nil
}

// Project methods
func (r *Repository) GetProjects() ([]models.Project, error) {
	query := `SELECT id, title, description, image, technologies, github, demo, featured, created_at, updated_at FROM projects ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("[GetProjects] Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		var desc, github, demo sql.NullString
		var image, technologies pq.StringArray
		err := rows.Scan(
			&p.ID, &p.Title, &desc, &image, &technologies, &github, &demo,
			&p.Featured, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			log.Printf("[GetProjects] Scan error: %v", err)
			return nil, err
		}
		p.Description = desc.String
		p.Image = models.StringArray(image)
		p.Technologies = models.StringArray(technologies)
		p.Github = github.String
		p.Demo = demo.String
		projects = append(projects, p)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetProjects] Rows error: %v", err)
		return nil, err
	}
	return projects, nil
}

// Skill methods
func (r *Repository) GetSkills() ([]models.Skill, error) {
	query := `SELECT id, name, level, icon_url, category FROM skills ORDER BY category, name`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("[GetSkills] Query error: %v", err)
		return nil, err
	}
	defer rows.Close()
	var skills []models.Skill
	for rows.Next() {
		var skill models.Skill
		var iconURL sql.NullString
		err := rows.Scan(&skill.ID, &skill.Name, &skill.Level, &iconURL, &skill.Category)
		if err != nil {
			log.Printf("[GetSkills] Scan error: %v", err)
			return nil, err
		}
		skill.IconURL = iconURL.String
		skills = append(skills, skill)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetSkills] Rows error: %v", err)
		return nil, err
	}
	return skills, nil
}

// Experience methods
func (r *Repository) GetExperiences() ([]models.Experience, error) {
	query := `SELECT id, company, position, location, start_date, end_date, current, description, logo, technologies FROM experiences ORDER BY start_date DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("[GetExperiences] Query error: %v", err)
		return nil, err
	}
	defer rows.Close()
	var experiences []models.Experience
	for rows.Next() {
		var exp models.Experience
		var location, description, logo sql.NullString
		var endDate sql.NullTime
		var technologies pq.StringArray
		err := rows.Scan(
			&exp.ID, &exp.Company, &exp.Position, &location, &exp.StartDate,
			&endDate, &exp.Current, &description, &logo, &technologies,
		)
		if err != nil {
			log.Printf("[GetExperiences] Scan error: %v", err)
			return nil, err
		}
		exp.Location = location.String
		exp.Description = description.String
		exp.Logo = logo.String
		if endDate.Valid {
			exp.EndDate = &endDate.Time
		}
		exp.Technologies = models.StringArray(technologies)
		experiences = append(experiences, exp)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetExperiences] Rows error: %v", err)
		return nil, err
	}
	return experiences, nil
}

// About Cards methods
func (r *Repository) GetAboutCards() ([]models.AboutCard, error) {
	query := `SELECT id, icon, title, description, gradient_colors, display_order, hover_rotate_y, created_at, updated_at FROM about_cards ORDER BY display_order`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("[GetAboutCards] Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var cards []models.AboutCard
	for rows.Next() {
		var card models.AboutCard
		var gradientColors pq.StringArray
		err := rows.Scan(
			&card.ID, &card.Icon, &card.Title, &card.Description,
			&gradientColors, &card.DisplayOrder, &card.HoverRotateY,
			&card.CreatedAt, &card.UpdatedAt,
		)
		if err != nil {
			log.Printf("[GetAboutCards] Scan error: %v", err)
			return nil, err
		}
		card.GradientColors = models.StringArray(gradientColors)
		cards = append(cards, card)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetAboutCards] Rows error: %v", err)
		return nil, err
	}
	return cards, nil
}

// Learning Items methods
func (r *Repository) GetLearningItems() ([]models.LearningItem, error) {
	query := `SELECT id, title, color, display_order, created_at, updated_at FROM learning_items ORDER BY display_order`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("[GetLearningItems] Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var items []models.LearningItem
	for rows.Next() {
		var item models.LearningItem
		err := rows.Scan(
			&item.ID, &item.Title, &item.Color, &item.DisplayOrder,
			&item.CreatedAt, &item.UpdatedAt,
		)
		if err != nil {
			log.Printf("[GetLearningItems] Scan error: %v", err)
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetLearningItems] Rows error: %v", err)
		return nil, err
	}
	return items, nil
}

// Goals methods
func (r *Repository) GetGoals() ([]models.Goal, error) {
	query := `SELECT id, title, color, display_order, created_at, updated_at FROM goals ORDER BY display_order`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("[GetGoals] Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var goals []models.Goal
	for rows.Next() {
		var goal models.Goal
		err := rows.Scan(
			&goal.ID, &goal.Title, &goal.Color, &goal.DisplayOrder,
			&goal.CreatedAt, &goal.UpdatedAt,
		)
		if err != nil {
			log.Printf("[GetGoals] Scan error: %v", err)
			return nil, err
		}
		goals = append(goals, goal)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetGoals] Rows error: %v", err)
		return nil, err
	}
	return goals, nil
}

// Fun Facts methods
func (r *Repository) GetFunFacts() ([]models.FunFact, error) {
	query := `SELECT id, emoji, text, rotate_angle, display_order, created_at, updated_at FROM fun_facts ORDER BY display_order`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("[GetFunFacts] Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var facts []models.FunFact
	for rows.Next() {
		var fact models.FunFact
		err := rows.Scan(
			&fact.ID, &fact.Emoji, &fact.Text, &fact.RotateAngle,
			&fact.DisplayOrder, &fact.CreatedAt, &fact.UpdatedAt,
		)
		if err != nil {
			log.Printf("[GetFunFacts] Scan error: %v", err)
			return nil, err
		}
		facts = append(facts, fact)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetFunFacts] Rows error: %v", err)
		return nil, err
	}
	return facts, nil
}
