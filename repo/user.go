package repo

import "database/sql"

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Photo      string `json:"photo"`
	IsVerified bool   `json:"is_verified"`
}

type UserRepo interface {
	CreateUser(user User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserPassword(email string) (string, error)
	UpdatePassword(email, HashedPassword string) (*User, error)
	UpdateUserProfile(user *User) (*User, error)
	UpdateUserVerification(email string, isVerified bool) (*User, error)
	DeleteUserByEmail(email string) error
}

type userRepo struct {
	dbCon *sql.DB
}

func NewUserRepo(dbCon *sql.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r *userRepo) CreateUser(user User) (*User, error) {

	query := `
		INSERT INTO users (first_name, last_name, phone, email, password, photo)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.dbCon.Exec(query, user.FirstName, user.LastName, user.Phone, user.Email, user.Password, user.Photo)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	return &user, nil
}

func (r *userRepo) GetUserByEmail(email string) (*User, error) {
	query := `
		SELECT id, first_name, last_name, phone, email, password, photo, is_verified
		FROM users
		WHERE email = ?
		LIMIT 1
	`

	user := &User{}
	err := r.dbCon.QueryRow(query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.Photo,
		&user.IsVerified,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no user found
		}
		return nil, err
	}

	return user, nil
}
func (r *userRepo) GetUserPassword(email string) (string, error) {
	query := `
		SELECT password
		FROM users
		WHERE email = ?
		LIMIT 1
	`

	var password string
	err := r.dbCon.QueryRow(query, email).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return password, nil
}

func (r *userRepo) UpdatePassword(email, HashedPassword string) (*User, error) {
	updateQuery := `
		UPDATE users
		SET password = ?
		WHERE email = ?
	`
	_, err := r.dbCon.Exec(updateQuery, HashedPassword, email)
	if err != nil {
		return nil, err
	}

	return r.GetUserByEmail(email)
}

func (r *userRepo) UpdateUserProfile(user *User) (*User, error) {
	query := `
		UPDATE users
		SET first_name = ?, last_name = ?, phone = ?
		WHERE email = ?
	`
	_, err := r.dbCon.Exec(query, user.FirstName, user.LastName, user.Phone, user.Email)
	if err != nil {
		return nil, err
	}

	return r.GetUserByEmail(user.Email)
}

func (r *userRepo) UpdateUserVerification(email string, isVerified bool) (*User, error) {
	query := `
		UPDATE users
		SET is_verified = ?
		WHERE email = ?
	`
	_, err := r.dbCon.Exec(query, isVerified, email)
	if err != nil {
		return nil, err
	}

	return r.GetUserByEmail(email)
}

func (r *userRepo) DeleteUserByEmail(email string) error {
	query := `DELETE FROM users WHERE email = ?`
	_, err := r.dbCon.Exec(query, email)
	return err
}
