package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/gvu0110/bookstore_utils-go/rest_errors"
)

func ResponseJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func ResponseRESTError(w http.ResponseWriter, err rest_errors.RESTError) {
	ResponseJSON(w, err.StatusCode(), err)
}
