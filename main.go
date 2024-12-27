package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/srisudarshanrg/go-fighter-jets/server/config"
	"github.com/srisudarshanrg/go-fighter-jets/server/database"
	"github.com/srisudarshanrg/go-fighter-jets/server/functions"
	"github.com/srisudarshanrg/go-fighter-jets/server/handlers"
	"github.com/srisudarshanrg/go-fighter-jets/server/models"
	"github.com/srisudarshanrg/go-fighter-jets/server/render"
	"github.com/srisudarshanrg/go-fighter-jets/server/validations"
)

const portNumber = ":2121"

var session *scs.SessionManager
var appConfig config.AppConfig

func main() {
	gob.Register(models.FighterJet{})

	// session
	session = scs.New()
	session.Cookie.Persist = true
	session.Lifetime = 24 * time.Hour
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// database handlers
	db, err := database.CreateDatabaseConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// template cache handlers
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Could not create template cache: ", err)
	}

	// app config handlers
	appConfig.TemplateCache = templateCache
	appConfig.ProjectCompleted = false
	appConfig.UseTemplateCache = appConfig.ProjectCompleted
	appConfig.Database = db
	appConfig.Session = session

	// handlers repository
	handlerRepo := handlers.HandlerRepository{
		AppConfig: appConfig,
	}
	handlers.RepositoryAccesshandlers(handlerRepo)

	// app config access
	functions.AppConfigAccessFunctions(appConfig)
	validations.AppConfigAccessValidations(appConfig)

	// routes
	server := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	log.Println("server running on port number", portNumber)
	server.ListenAndServe()
}

func routes() http.Handler {
	mux := chi.NewRouter()

	// middleware
	mux.Use(SessionLoadAndSave)

	// routes for get requests
	mux.Get("/", handlers.Repository.Home)
	mux.Get("/fighter-category", handlers.Repository.FighterCategory)
	mux.Get("/compare", handlers.Repository.Compare)
	mux.Get("/jet-details", handlers.Repository.JetDetails)

	// routes for post requests
	mux.Post("/", handlers.Repository.HomePost)
	mux.Post("/compare", handlers.Repository.ComparePost)

	// routes for static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func SessionLoadAndSave(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
