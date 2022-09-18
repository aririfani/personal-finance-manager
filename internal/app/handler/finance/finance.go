package finance

import (
	"encoding/json"
	"errors"
	"github.com/aririfani/personal-finance-manager/internal/app/service"
	"github.com/aririfani/personal-finance-manager/internal/app/service/finance"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Handler interface {
	CreateFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
	GetAllFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
	UpdateFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
	GetFinanceByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
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

// CreateFinance ...
func (h *handler) CreateFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request finance.Finance
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (finance.Finance{}) {
		err = errors.New("error payload")
		return
	}

	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)
	request.UserID = int64(userID)
	returnData, err = h.service.Finance.CreateFinance(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

// GetAllFinance ...
func (h *handler) GetAllFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)

	request := finance.GetAllFinanceReq{
		Page:   page,
		Limit:  limit,
		UserID: int64(userID),
	}

	returnData, err = h.service.Finance.GetAllFinance(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

// UpdateFinance ...
func (h *handler) UpdateFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request finance.Finance
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (finance.Finance{}) {
		err = errors.New("error payload")
		return
	}

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)
	request.UserID = int64(userID)
	request.ID = id

	returnData, err = h.service.Finance.UpdateFinance(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

// GetFinanceByID ...
func (h *handler) GetFinanceByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	returnData, err = h.service.Finance.GetFinanceByID(ctx, id)
	
	w.Header().Add("Content-Type", "application/json")

	return
}
