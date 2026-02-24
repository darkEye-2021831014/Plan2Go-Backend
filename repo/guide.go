package repo

import (
	"database/sql"
	"fmt"
	"time"
)

type Guide struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	City              string    `json:"city"`
	HourlyFee         float64   `json:"hourly_fee"`
	Languages         string    `json:"languages"`
	YearsOfExperience int       `json:"years_of_experience"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type GuideRepo interface {
	GetGuides(city string) ([]Guide, error)
	CreateGuide(g Guide) (*Guide, error)
}

type guideRepo struct {
	dbCon *sql.DB
}

func NewGuideRepo(dbCon *sql.DB) GuideRepo {
	return &guideRepo{dbCon: dbCon}
}

func (r *guideRepo) GetGuides(city string) ([]Guide, error) {
	query := `
		SELECT g.id, g.user_id, u.first_name, u.last_name, g.city, g.hourly_fee,
		       g.languages, g.years_of_experience
		FROM guides g
		JOIN users u ON g.user_id = u.id
		WHERE 1=1
	`
	args := []interface{}{}

	if city != "" {
		query += " AND g.city = ?"
		args = append(args, city)
	}

	rows, err := r.dbCon.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guides []Guide
	for rows.Next() {
		var g Guide
		err := rows.Scan(
			&g.ID, &g.UserID, &g.FirstName, &g.LastName,
			&g.City, &g.HourlyFee, &g.Languages, &g.YearsOfExperience,
		)
		if err != nil {
			return nil, err
		}
		guides = append(guides, g)
	}

	return guides, nil
}

func (r *guideRepo) CreateGuide(g Guide) (*Guide, error) {

	var exists bool
	err := r.dbCon.QueryRow("SELECT EXISTS(SELECT 1 FROM guides WHERE user_id = ?)", g.UserID).Scan(&exists)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("user with id %d is already registered as a guide", g.UserID)
	}

	// Insert new guide
	query := `
        INSERT INTO guides (user_id, city, hourly_fee, languages, years_of_experience, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, NOW(), NOW())
    `
	result, err := r.dbCon.Exec(query, g.UserID, g.City, g.HourlyFee, g.Languages, g.YearsOfExperience)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	g.ID = int(id)
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()

	return &g, nil
}
