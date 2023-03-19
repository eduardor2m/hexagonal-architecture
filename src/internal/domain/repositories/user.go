package repositories

import (
	"database/sql"

	"github.com/eduardor2m/hexagonal-architecture/src/internal/domain/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Save(user *models.User) error {
	_, err := u.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := u.db.QueryRow("SELECT name, email FROM users WHERE email = ?", email).Scan(&user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
