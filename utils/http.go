package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/gorilla/mux"
	"github.com/techx/portal/constants"
)

// ToDo: Refactor this handling of auth subject and unique id. Move to appropriate package.

func GetAuthSubject(r *http.Request) string {
	apiName := mux.CurrentRoute(r).GetName()
	if apiName == constants.APINameUserRegister {
		var req struct {
			WorkEmail string `json:"work_email"`
		}
		decoder := json.NewDecoder(r.Body)
		_ = decoder.Decode(&req)

		return req.WorkEmail
	}

	return r.Header.Get(constants.HeaderXUserUUID)
}

func GetUniqueRequestID(r *http.Request) string {
	uniqueID := r.Header.Get(constants.HeaderXUserUUID)
	apiName := mux.CurrentRoute(r).GetName()
	if slices.Contains(constants.AuthRoutes, apiName) {
		var req struct {
			Value string `json:"value"`
		}
		decoder := json.NewDecoder(r.Body)
		_ = decoder.Decode(&req)

		uniqueID = req.Value
	}

	if slices.Contains(constants.AdminRoutes, apiName) {
		uniqueID = r.Header.Get(constants.HeaderClientID)
	}

	if apiName == constants.APINameUserRegister {
		var req struct {
			PhoneNumber string `json:"phone_number"`
		}
		decoder := json.NewDecoder(r.Body)
		_ = decoder.Decode(&req)

		uniqueID = req.PhoneNumber
	}

	return uniqueID
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	body, err := json.Marshal(v)
	if err != nil {
		msg := fmt.Sprintf("Could not convert the given object to JSON: %v", err)
		writeText(w, http.StatusInternalServerError, msg)

		return
	}

	w.Header().Set(constants.HeaderContentType, constants.ApplicationJSON)

	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func writeText(w http.ResponseWriter, status int, body string) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(body))
}
