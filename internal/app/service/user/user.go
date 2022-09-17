package user

import (
	"context"
	"errors"
	"github.com/aririfani/personal-finance-manager/internal/app/repository"
	"github.com/aririfani/personal-finance-manager/internal/app/repository/user"
	pkgError "github.com/aririfani/personal-finance-manager/internal/pkg/errors"
	"github.com/aririfani/personal-finance-manager/internal/pkg/token"
	"github.com/aririfani/personal-finance-manager/internal/pkg/utils"
	"github.com/ulule/deepcopier"
	"time"
)

// srv ...
type srv struct {
	repo  *repository.Repositories
	token token.IToken
}

// NewSrv ..
func NewSrv(repo *repository.Repositories, token token.IToken) Service {
	return &srv{
		repo:  repo,
		token: token,
	}
}

// CreateUser ...
func (s *srv) CreateUser(ctx context.Context, param User) (returnData User, err error) {
	userRepo := user.User{}

	// encrypt password
	password, err := utils.EncryptString(param.Password)
	if err != nil {
		return
	}

	param.Password = password
	_ = deepcopier.Copy(param).To(&userRepo)

	res, err := s.repo.User.CreateUser(ctx, userRepo)
	if err != nil {
		return
	}
	_ = deepcopier.Copy(res).To(&returnData)

	return
}

func (s *srv) Login(ctx context.Context, param User) (returnData LoginRes, err error) {
	var userRepo user.User
	_ = deepcopier.Copy(param).To(&userRepo)

	res, err := s.repo.User.GetUserByEmail(ctx, userRepo)
	if err != nil {
		return
	}

	validPassword, err := utils.CompareStringValid(param.Password, res.Password)
	if !validPassword {
		err = &pkgError.BadRequest{Err: errors.New("email or password invalid")}
		return
	}

	var getTokenRes token.GetToken
	var accessTokenExpiresAt int64
	var expiredAt uint64

	getTokenRes.AccessToken, accessTokenExpiresAt, err = s.token.GetImplicitToken(
		token.Payload{
			ID:       res.ID,
			Email:    res.Email,
			Phone:    res.Phone,
			FullName: res.FullName,
		}, expiredAt)

	getTokenRes.AccessTokenExpiresAt = time.Unix(accessTokenExpiresAt, 0)

	returnData = LoginRes{
		AccessToken: getTokenRes.AccessToken,
		ExpiredAt:   getTokenRes.AccessTokenExpiresAt,
	}

	return
}
