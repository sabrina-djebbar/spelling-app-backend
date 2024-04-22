package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure/repo"
	"time"
)

type Repository struct {
	q repo.Queries
}

func NewRepo(queries repo.Queries) Repository {
	return Repository{q: queries}
}

func (r *Repository) GetUser(ctx context.Context, id string) (*repo.User, error) {
	u, err := r.q.GetUser(ctx, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return &u, nil
}

type CreateUserParams struct {
	ID          string
	Username    string
	DateOfBirth time.Time
	ParentCode  string
	Password    string
}

func (r *Repository) CreateUser(ctx context.Context, user CreateUserParams) (*repo.User, error) {
	u, err := r.q.CreateUser(ctx, repo.CreateUserParams{ID: user.ID, Username: user.Username, DateOfBirth: sql.NullTime{Time: user.DateOfBirth}, ParentCode: user.ParentCode})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	err = r.q.CreateCredentials(ctx, repo.CreateCredentialsParams{
		UserID: user.ID,
		Crypt:  user.Password,
	})
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) ListUsers(ctx context.Context) (*[]repo.User, error) {
	u, err := r.q.ListUsers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return &u, nil
}
