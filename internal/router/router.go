package router

import (
	"net/http"
	"path/filepath"

	"github.com/melnikov-anton/workday-api/internal/config"
	"github.com/melnikov-anton/workday-api/internal/handlers"
)

var appConfig *config.AppConfig

func NewRouter(app *config.AppConfig) *http.ServeMux {
	appConfig = app

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomePage)
	mux.HandleFunc("/api", handlers.InfoApi)
	mux.HandleFunc("/api/", handlers.InfoApi)

	staticDir := filepath.Join(appConfig.AppRootDir, "static")
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
