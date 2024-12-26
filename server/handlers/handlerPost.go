package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/srisudarshanrg/go-fighter-jets/server/functions"
	"github.com/srisudarshanrg/go-fighter-jets/server/models"
	"github.com/srisudarshanrg/go-fighter-jets/server/render"
)

// HomePost handles post requests to the home page
func (app *HandlerRepository) HomePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	jetDetailID := r.Form.Get("jetDetailID")
	searchKey := r.Form.Get("searchKey")

	if jetDetailID != "" {
		id, err := strconv.Atoi(jetDetailID)
		if err != nil {
			log.Println(err)
		}
		jet, err := functions.GetFighterJetByID(id)
		if err != nil {
			log.Println(err)
		}
		app.AppConfig.Session.Put(r.Context(), "jetDetails", jet)

		http.Redirect(w, r, "/jet-details", http.StatusSeeOther)
	} else if searchKey != "" {
		results, err := functions.GetFighterJetsBySearchKey(searchKey)
		if err != nil {
			log.Println(err)
		}

		postData := map[string]interface{}{}
		postData["searchResults"] = results
		postData["lengthSearchResults"] = len(results)
		err = render.RenderTemplate(w, r, "home.page.tmpl", models.TemplateData{
			PostData: postData,
			Data:     data,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
