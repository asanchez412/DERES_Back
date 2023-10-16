package user

import (
	"context"
	"database/sql"

	"github.com/andsanchez/DERES_Back/internal/domain"
)

// Repository encapsulates the storage of a section.
type Repository interface {
	Get(ctx context.Context, username, password string) (domain.User, error)
	Create(ctx context.Context, user domain.User) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// This constants are the queries that the repo entity will use
const (
	GetUser    string = "SELECT * FROM users WHERE user_name = ? AND password = ?;"
	CreateUser string = "INSERT INTO users (user_name, password) VALUES (?, ?);"
)

func (r *repository) Get(ctx context.Context, username, password string) (domain.User, error) {
	row := r.db.QueryRowContext(ctx, GetUser, username, password)
	u := domain.User{}
	err := row.Scan(&u.Username, &u.Password)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r *repository) Create(ctx context.Context, user domain.User) error {
	stmt, err := r.db.PrepareContext(ctx, CreateUser)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, &user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}
