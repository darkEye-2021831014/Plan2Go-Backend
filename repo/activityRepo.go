package repo

import (
	"database/sql"
	"fmt"
	"plan2go-backend/models"
)

type ActivityRepo struct {
	DB *sql.DB
}

func NewActivityRepo(db *sql.DB) *ActivityRepo {
	return &ActivityRepo{DB: db}
}

func (r *ActivityRepo) CreateActivity(a models.Activity) error {
	query := `
        INSERT INTO user_activity (user_id, activity_type, description, page)
        VALUES (?, ?, ?, ?)
    `
	_, err := r.DB.Exec(query, a.UserID, a.Action, a.Description, a.Page)
	if err != nil {
		fmt.Println("CreateActivity ERROR:", err)
	}
	return err
}

func (r *ActivityRepo) QueryActivity(userID int) ([]models.Activity, error) {
	query := `
        SELECT 
            id,
            user_id,
            activity_type,
            description,
            IFNULL(page, '')
        FROM user_activity
        WHERE user_id = ?
        ORDER BY created_at DESC
    `

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		fmt.Println("QueryActivity ERROR:", err)
		return nil, err
	}
	defer rows.Close()

	var activities []models.Activity

	for rows.Next() {
		var a models.Activity
		if err := rows.Scan(
			&a.ID,
			&a.UserID,
			&a.Action,
			&a.Description,
			&a.Page,
		); err != nil {
			fmt.Println("Scan ERROR:", err)
			return nil, err
		}
		activities = append(activities, a)
	}

	return activities, nil
}
