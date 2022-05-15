package handlers

import (
	"fmt"
	"net/http"

	"github.com/melnikov-anton/workday-api/internal/config"
)

var appConfig *config.AppConfig

const HEADER_CONTENT_TYPE = "Content-Type"

// InitHandlers makes app config available
func InitHandlers(app *config.AppConfig) {
	appConfig = app
}

// sendJsonResponse sends json response as it is
func sendJsonResponse(rw http.ResponseWriter, body []byte, code int) {
	rw.Header().Set(HEADER_CONTENT_TYPE, "application/json")
	rw.WriteHeader(code)
	rw.Write(body)
}

// sendErrorJsonResponse sends json response with error key
func sendErrorJsonResponse(rw http.ResponseWriter, body []byte, code int) {
	resp := fmt.Sprintf("{\"error\":\"%s\"}", body)
	rw.Header().Set(HEADER_CONTENT_TYPE, "application/json")
	rw.WriteHeader(code)
	rw.Write([]byte(resp))
}

// sendSimpleResponse sends simple text response
func sendSimpleResponse(rw http.ResponseWriter, body []byte, code int) {
	rw.Header().Set(HEADER_CONTENT_TYPE, "text/plain")
	rw.WriteHeader(code)
	rw.Write(body)
}
