package user

import "context"

// Service ...
type Service interface {
	// CreateUser ...
	CreateUser(ctx context.Context, param User) (returnData User, err error)

	// Login ...
	Login(ctx context.Context, param User) (returnData LoginRes, err error)

	// Profile ...
	Profile(ctx context.Context, id int64) (returnData Res, err error)
}
