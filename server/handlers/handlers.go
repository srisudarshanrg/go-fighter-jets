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
	jets, err := functions.GetAllFighterJets()
	if err != nil {
		log.Println(err)
	}

	data["jets"] = jets
	err = render.RenderTemplate(w, r, "compare.page.tmpl", models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

// JetDetails is the handler for the jet details page
func (app *HandlerRepository) JetDetails(w http.ResponseWriter, r *http.Request) {
	jetInterface := app.AppConfig.Session.Get(r.Context(), "jetDetails")
	jet, ok := jetInterface.(models.FighterJet)
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	link := app.AppConfig.Session.Get(r.Context(), "link").(string)

	data["jetDetail"] = jet
	data["link"] = link
	err := render.RenderTemplate(w, r, "jet-details.page.tmpl", models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}

	app.AppConfig.Session.Remove(r.Context(), "jetDetails")
	app.AppConfig.Session.Remove(r.Context(), "link")
}
