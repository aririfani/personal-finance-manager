package user

import "context"

// Repository ...
type Repository interface {
	// CreateUser ...
	CreateUser(ctx context.Context, user User) (returnData User, err error)

	// GetUserByEmail ...
	GetUserByEmail(ctx context.Context, user User) (returnData User, err error)
}
