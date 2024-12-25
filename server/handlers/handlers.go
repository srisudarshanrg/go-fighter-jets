package handlers

import (
	"log"
	"net/http"

	"github.com/srisudarshanrg/go-fighter-jets/server/functions"
	"github.com/srisudarshanrg/go-fighter-jets/server/models"
	"github.com/srisudarshanrg/go-fighter-jets/server/render"
)

var data = make(map[string]interface{})

// Home is the handler for the home page
func (app *HandlerRepository) Home(w http.ResponseWriter, r *http.Request) {
	jets, err := functions.GetAllFighterJets()
	if err != nil {
		log.Println(err)
	}

	data["jets"] = jets
	err = render.RenderTemplate(w, r, "home.page.tmpl", models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

// FighterCategory is the handler for the fighter category page
func (app *HandlerRepository) FighterCategory(w http.ResponseWriter, r *http.Request) {
	err := render.RenderTemplate(w, r, "fighter-category.page.tmpl", models.TemplateData{})
	if err != nil {
		log.Println(err)
	}
}

// Compare is the handler for the compare page
func (app *HandlerRepository) Compare(w http.ResponseWriter, r *http.Request) {
	err := render.RenderTemplate(w, r, "compare.page.tmpl", models.TemplateData{})
	if err != nil {
		log.Println(err)
	}
}
