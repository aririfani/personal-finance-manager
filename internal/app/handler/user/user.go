package user

import (
	"encoding/json"
	"errors"
	"github.com/aririfani/personal-finance-manager/internal/app/service"
	"github.com/aririfani/personal-finance-manager/internal/app/service/user"
	"io/ioutil"
	"net/http"
)

// Handler ...
type Handler interface {
	// CreateUser ...
	CreateUser(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
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
