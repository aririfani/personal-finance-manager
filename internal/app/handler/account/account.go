package account

import (
	"encoding/json"
	"errors"
	"github.com/aririfani/personal-finance-manager/internal/app/service"
	"github.com/aririfani/personal-finance-manager/internal/app/service/account"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Handler interface {
	// CreateAccount ...
	CreateAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// UpdateAccount ...
	UpdateAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// DeleteAccount ...
	DeleteAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// GetAllAccount ...
	GetAllAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// GetByAccountByID ...
	GetByAccountByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
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

// CreateAccount ...
func (h *handler) CreateAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request account.Account
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (account.Account{}) {
		err = errors.New("error payload")
		return
	}

	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)
	request.UserID = int64(userID)
	returnData, err = h.service.Account.CreateAccount(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

// UpdateAccount ...
func (h *handler) UpdateAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request account.Account
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (account.Account{}) {
		err = errors.New("error payload")
		return
	}

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	request.ID = id
	returnData, err = h.service.Account.UpdateAccount(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

// DeleteAccount ...
func (h *handler) DeleteAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	returnData, err = h.service.Account.DeleteAccount(ctx, id)
	w.Header().Add("Content-Type", "application/json")
	return
}

// GetAllAccount ...
func (h *handler) GetAllAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)

	request := account.GetAllAccountReq{
		Page:   page,
		Limit:  limit,
		UserID: int64(userID),
	}

	returnData, err = h.service.Account.GetAllAccount(ctx, request)
	w.Header().Add("Content-Type", "application/json")
	return
}

// GetByAccountByID ...
func (h *handler) GetByAccountByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	returnData, err = h.service.Account.GetAccountByID(ctx, id)
	w.Header().Add("Content-Type", "application/json")

	return
}
