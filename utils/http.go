package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/techx/portal/constants"
)

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
