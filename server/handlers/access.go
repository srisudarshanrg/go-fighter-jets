package handlers

import "github.com/srisudarshanrg/go-fighter-jets/server/config"

var AppConfig HandlerRepository

// HandlerRepository is the repository for the handlers package which contains the app config, database and session
type HandlerRepository struct {
	AppConfig config.AppConfig
}

// RepositoryAccesshandlers provides the handlers package with access to the repository
func RepositoryAccesshandlers(configSet HandlerRepository) {
	AppConfig = configSet
}
