// Package httputils - Provide support functions to help http integration flow
package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/eviccari/rest-http-utils/httperrors"
)

const (
	CONTENT_TYPE_STRING     = "Content-type"
	APPLICATION_JSON_STRING = "application/json"
)

type DefaultHttpJSONResponse struct {
	Message        string `json:"message"`
	HTTPStatusCode uint   `json:"statusCode"`
}

// NewDefaultHttpJSONResponse - Create new DefaultHttpJSONResponse instance
func NewDefaultHttpJSONResponse(message string, statusCode uint) *DefaultHttpJSONResponse {
	return &DefaultHttpJSONResponse{
		Message:        message,
		HTTPStatusCode: statusCode,
	}
}

// NewDefaultSuccessJSONResponse - Create new DefaultHttpJSONResponse instance with OK (200) http status
func NewDefaultSuccessJSONResponse(message string) *DefaultHttpJSONResponse {
	return &DefaultHttpJSONResponse{
		Message:        message,
		HTTPStatusCode: http.StatusOK,
	}
}

// WriteJSONResponse - Write JSON payload into ResponseWriter
func WriteJSONResponse(w http.ResponseWriter, payload interface{}, statusCode int) {
	response, err := json.Marshal(payload)
	if err != nil {
		WriteJSONErrorResponse(w, httperrors.NewInternalServerError(err))
	} else {
		w.Header().Add(CONTENT_TYPE_STRING, APPLICATION_JSON_STRING)
		w.WriteHeader(statusCode)
		w.Write(response)
	}
}

// WriteJSONErrorResponse - Write JSON Error payload into ResponseWriter
func WriteJSONErrorResponse(w http.ResponseWriter, errorMessage httperrors.IHTTPError) {
	response, _ := json.Marshal(errorMessage)
	w.Header().Add(CONTENT_TYPE_STRING, APPLICATION_JSON_STRING)
	w.WriteHeader(errorMessage.StatusCode())
	w.Write(response)
}

// GetValueFromHeader - Get value from specific header
func GetValueFromHeader(r *http.Request, key string) string {
	values := r.Header[http.CanonicalHeaderKey(key)]
	if len(values) > 0 {
		return values[0]
	}

	return ""
}
