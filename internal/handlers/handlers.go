package handlers

import "github.com/melnikov-anton/workday-api/internal/config"

var appConfig *config.AppConfig

func InitHandlers(app *config.AppConfig) {
	appConfig = app
}
