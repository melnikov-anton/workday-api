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

	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
