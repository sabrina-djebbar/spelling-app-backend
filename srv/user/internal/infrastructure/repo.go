package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
  
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/id"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/serr"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure/repo"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
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
			return nil, serr.New("No user found with given ID")
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

		return nil, serr.Wrap(err, serr.WithMessage("Unable to create user"))

	}

	err = r.q.CreateCredentials(ctx, repo.CreateCredentialsParams{
		ID:     id.Generate("credential"),
		UserID: u.ID,
		Crypt:  user.Password,
	})
	if err != nil {
		return nil, serr.Wrap(err, serr.WithMessage("Unable to create credentials"))
	}

	return &u, nil
}

func (r *Repository) ListUsers(ctx context.Context) (*[]repo.User, error) {
	u, err := r.q.ListUsers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, serr.Wrap(err, serr.WithMessage("Unable to list users"))

	}
	return &u, nil
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (*repo.User, error) {
	u, err := r.q.FindByUsername(ctx, username)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, serr.Wrap(err, serr.WithMessage("No user found with given username: "+username))
		}
		return nil, serr.Wrap(err, serr.WithMessage("Unable to find user with given username: "+username))
	}
	return &u, nil
}

type FindCredentialParams struct {
	UserID string
	Crypt  string
}

func (r *Repository) FindCredentials(ctx context.Context, credentials FindCredentialParams) (string, error) {
	uid, err := r.q.FindCredentials(ctx, repo.FindCredentialsParams{UserID: credentials.UserID, Crypt: credentials.Crypt})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", serr.Wrap(err, serr.WithMessage("No credentials found"))
		}
		return "", serr.Wrap(err, serr.WithMessage("Unable to credentials"))
	}
	return uid, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	panic("implement me")
}

func (r *Repository) EditParentCode(ctx context.Context, id string, code string) (*repo.User, error) {
	u, err := r.q.UpdateParentCode(ctx, repo.UpdateParentCodeParams{ID: id, ParentCode: code})
	if err != nil {
		return nil, serr.Wrap(err, serr.WithMessage("Unable to update parent code"), serr.WithCode(serr.ErrCodeInternalService))
	}
	return &u, nil
}

func (r *Repository) EditUser(ctx context.Context, args models.User) (*repo.User, error) {

	user, err := r.q.UpdateUser(ctx, repo.UpdateUserParams{
		ID:          args.ID,
		DateOfBirth: database.TimeToSQLNullTime(args.DateOfBirth),
		Username:    args.Username,
	})
	if err != nil {
		return nil, serr.Wrap(err, serr.WithMessage("Unable to update parent code"), serr.WithCode(serr.ErrCodeInternalService))
	}
	return &user, nil
}
