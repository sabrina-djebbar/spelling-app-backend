package infrastructure

import (
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure/repo"
)

type repository struct {
	q repo.Querier
}

func NewRepo(queries sql.Querier) repo {
	return repo{
		q: queries,
	}
}
func (r *repository) GetUser(ctx context.Context, id string) (*User, error) {
	u, err := r.q.GetUser(ctx, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	if u == nil {
		return nil, nil
	}
	return &u, nil
}
