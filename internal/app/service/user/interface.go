package user

import "context"

// Service ...
type Service interface {
	// CreateUser ...
	CreateUser(ctx context.Context, param User) (returnData User, err error)

	// Login ...
	Login(ctx context.Context, param User) (returnData LoginRes, err error)
}
