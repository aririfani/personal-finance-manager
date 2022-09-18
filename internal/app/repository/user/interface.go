package user

import "context"

// Repository ...
type Repository interface {
	// CreateUser ...
	CreateUser(ctx context.Context, user User) (returnData User, err error)

	// GetUserByEmail ...
	GetUserByEmail(ctx context.Context, user User) (returnData User, err error)

	// GetUserByID ...
	GetUserByID(ctx context.Context, id int64) (returnData User, err error)
}
