package user

import "context"

type repo struct {
	DB Repository
}

func NewRepo(db Repository) Repository {
	return &repo{
		db,
	}
}

// CreateUser ...
func (r *repo) CreateUser(ctx context.Context, user User) (returnData User, err error) {
	returnData, err = r.DB.CreateUser(ctx, user)
	if err != nil {
		return
	}

	return
}

// GetUserByEmail ...
func (r *repo) GetUserByEmail(ctx context.Context, user User) (returnData User, err error) {
	returnData, err = r.DB.GetUserByEmail(ctx, user)
	if err != nil {
		return
	}

	return
}
