package repository

import (
	"context"

	"github.com/fajrimgfr/auth-notes-saas-backend/domain"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	database   *sqlx.DB
	collection string
}

func NewUserRepository(db *sqlx.DB, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	_, err := ur.database.NamedExecContext(c, `INSERT INTO users (id, email, password_hash, name, created_at, updated_at) VALUES (:id, :email, :password_hash, :name, :created_at, :updated_at)`, user)
	return err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := ur.database.GetContext(c, &user, `SELECT id, email, password_hash, name, created_at, updated_at FROM users WHERE email = $1`, email)
	return user, err
}

// func (ur *userRepository) Fetch(c context.Context, user *domain.User) error {
// 	return c.Err()
// }

// func (ur *userRepository) GetByEmail(c context.Context, user *domain.User) error {
// 	return c.Err()
// }

// func (ur *userRepository) GetByID(c context.Context, user *domain.User) error {
// 	return c.Err()
// }
