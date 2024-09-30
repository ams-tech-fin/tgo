package repository

import (
	"context"
	"tgo/api/internal/domain"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User = domain.User

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, name string, email string, password string) (User, error) {
	var user User
	query := `INSERT INTO users (name,email,password) VALUES ($1,$2,$3) RETURNING id, name,email`
	err := r.db.QueryRowxContext(ctx, query, name, email, password).StructScan(&user)
	return user, err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	var user User
	query := `SELECT id, name,email FROM users WHERE id = $1`
	err := r.db.GetContext(ctx, &user, query, id)
	return user, err
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User
	query := `SELECT id, name, email FROM users`
	err := r.db.SelectContext(context.Background(), &users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
