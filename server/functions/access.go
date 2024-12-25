package functions

import (
	"database/sql"

	"github.com/srisudarshanrg/go-fighter-jets/server/config"
)

var appConfig config.AppConfig
var db *sql.DB

// AppConfigAccessFunctions provides the functions package with access to the app config
func AppConfigAccessFunctions(a config.AppConfig) {
	appConfig = a
	db = appConfig.Database
}
