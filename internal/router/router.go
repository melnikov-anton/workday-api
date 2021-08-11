package router

import (
	"net/http"

	"github.com/melnikov-anton/workday-api/internal/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomePage)
	mux.HandleFunc("/api", handlers.InfoApi)
	mux.HandleFunc("/api/", handlers.InfoApi)

	return mux
}
