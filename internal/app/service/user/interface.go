package user

import "context"

// Service ...
type Service interface {
	// CreateUser ...
	CreateUser(ctx context.Context, param User) (returnData User, err error)
}
