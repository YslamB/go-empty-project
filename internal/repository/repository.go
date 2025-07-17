package repository

import (
	"context"

	"empty/internal/model"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.UserCrete) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) returning id",
		user.Name, user.Email, user.Password).Scan(&id)
	return id, err
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return &user, err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	err := pgxscan.Select(
		ctx, r.db,
		&users, "SELECT * FROM users",
	)

	if err != nil {
		return nil, err
	}

	return users, nil
}
