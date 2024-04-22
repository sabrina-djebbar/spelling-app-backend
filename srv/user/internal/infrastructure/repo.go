package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/id"
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
	Username    string
	DateOfBirth time.Time
	ParentCode  string
	Password    string
}

func (r *Repository) CreateUser(ctx context.Context, user CreateUserParams) (*repo.User, error) {
	req := repo.CreateUserParams{
		ID:          id.Generate("user"),
		Username:    user.Username,
		DateOfBirth: sql.NullTime{Time: user.DateOfBirth},
		ParentCode:  user.ParentCode,
	}

	u, err := r.q.CreateUser(ctx, req)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("unable to create user: " + err.Error())
		}
		return nil, errors.New("other user error\n " + err.Error())
	}

	err = r.q.CreateCredentials(ctx, repo.CreateCredentialsParams{
		ID:     id.Generate("credential"),
		UserID: u.ID,
		Crypt:  user.Password,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("unable to create credentials: " + err.Error())
		}
		return nil, errors.New("other credentials error\n " + err.Error())
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

func (r *Repository) FindByUsername(ctx context.Context, username string) (*repo.User, error) {
	u, err := r.q.FindByUsername(ctx, username)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("unable to find user by username: " + username)
		}
		return nil, err
	}
	return &u, nil
}

type FindCredentialParams struct {
	UserID string
	Crypt  string
}

func (r *Repository) FindCredentials(ctx context.Context, credentials FindCredentialParams) (string, error) {
	id, err := r.q.FindCredentials(ctx, repo.FindCredentialsParams{UserID: credentials.UserID, Crypt: credentials.Crypt})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			panic(errors.New("unable to find credentials for user: " + credentials.UserID))
		}
		panic(err)
	}
	return id, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	panic("implement me")
}
