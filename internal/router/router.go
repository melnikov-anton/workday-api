package router

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/melnikov-anton/workday-api/internal/config"
	"github.com/melnikov-anton/workday-api/internal/handlers"
)

var appConfig *config.AppConfig

func NewRouter(app *config.AppConfig) *mux.Router {
	appConfig = app

	mux := mux.NewRouter()

	mux.HandleFunc("/", handlers.HomePage)
	mux.HandleFunc("/api", handlers.InfoApi)
	mux.HandleFunc("/api/", handlers.InfoApi)
	mux.HandleFunc("/api/{cc}/workday/today", handlers.WorkdayToday)

	staticDir := filepath.Join(appConfig.AppRootDir, "static")
	fs := http.FileServer(http.Dir(staticDir))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return mux
}
