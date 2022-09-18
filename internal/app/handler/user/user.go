package user

import (
	"encoding/json"
	"errors"
	"github.com/aririfani/personal-finance-manager/internal/app/service"
	"github.com/aririfani/personal-finance-manager/internal/app/service/user"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"net/http"
)

// Handler ...
type Handler interface {
	// CreateUser ...
	CreateUser(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	//Login ...
	Login(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// Profile ...
	Profile(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
}

type handler struct {
	service *service.Services
}

// NewHandler ...
func NewHandler(services *service.Services) Handler {
	return &handler{
		service: services,
	}
}

// CreateUser ...
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request user.User
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (user.User{}) {
		err = errors.New("error payload")
		return
	}

	returnData, err = h.service.User.CreateUser(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request user.User
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (user.User{}) {
		err = errors.New("error payload")
		return
	}

	returnData, err = h.service.User.Login(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

func (h *handler) Profile(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)

	returnData, err = h.service.User.Profile(ctx, int64(userID))
	w.Header().Add("Content-Type", "application/json")
	return
}
