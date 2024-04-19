package rpc

import "context"

func GetUser(ctx context.Context, req GetUserRequest) (*User, error) {
	userReq := repo.GetUserParams{ID: req.UserId}
	res, err := r.groupRepo.GetGroupByID(ctx, req.GroupID)
	if err != nil {
		return nil, err
	}

	return &client.GetGroupResponse{Group: *res}, nil
	queries := repo.New(db)
	user, err = r.repo.GetUser(ctx, userReq)
	if err != nil {
		return nil, error
	}
	return user, err
}
