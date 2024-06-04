package request

// Path: handler/request/get_user_request.go

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/techx/portal/errors"
)

type GetUserByIDRequest struct {
	UserID int64 `json:"-"`
}

func NewGetUserByIDRequest(r *http.Request) (*GetUserByIDRequest, error) {
	var req GetUserByIDRequest
	var err error

	vars := mux.Vars(r)
	userIDStr := vars["userID"]
	req.UserID, err = strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (g GetUserByIDRequest) Validate() error {
	if g.UserID <= 0 {
		return errors.New("invalid user id")
	}

	return nil
}
