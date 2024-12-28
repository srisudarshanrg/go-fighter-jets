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
		app.AppConfig.Session.Put(r.Context(), "link", "/")

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

func (app *HandlerRepository) FighterCategoryPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	generationName := r.Form.Get("generationName")
	centuryName := r.Form.Get("centuryName")
	countryName := r.Form.Get("countryName")

	if generationName != "" {
		generationNameFloat, err := strconv.ParseFloat(generationName, 64)
		if err != nil {
			log.Println(err)
		}

		jets, err := functions.GetJetsByGeneration(generationNameFloat)
		if err != nil {
			log.Println(err)
		}

		postData := map[string]interface{}{}
		postData["generationJets"] = jets
		err = render.RenderTemplate(w, r, "fighter-category-results.page.tmpl", models.TemplateData{
			PostData: postData,
		})
		if err != nil {
			log.Println(err)
		}
	} else if centuryName != "" {
		var jets []models.FighterJet
		if centuryName == "1900" {
			jets, _, err = functions.GetJetsByCentury()
			if err != nil {
				log.Println(err)
			}
		} else {
			_, jets, err = functions.GetJetsByCentury()
			if err != nil {
				log.Println(err)
			}
		}
		postData := map[string]interface{}{}
		postData["centuryJets"] = jets
		err = render.RenderTemplate(w, r, "fighter-category-results.page.tmpl", models.TemplateData{
			PostData: postData,
		})
		if err != nil {
			log.Println(err)
		}
	} else if countryName != "" {
		jets, err := functions.GetJetsByCountry(countryName)
		if err != nil {
			log.Println(err)
		}

		postData := map[string]interface{}{}
		postData["countryJets"] = jets
		err = render.RenderTemplate(w, r, "fighter-category-results.page.tmpl", models.TemplateData{
			PostData: postData,
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func (app *HandlerRepository) ComparePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	fighterJet1 := r.Form.Get("fighterJet1")
	fighterJet2 := r.Form.Get("fighterJet2")

	jetDetailID := r.Form.Get("jetDetailID")

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
		app.AppConfig.Session.Put(r.Context(), "link", "/compare")

		http.Redirect(w, r, "/jet-details", http.StatusSeeOther)
	} else {
		if fighterJet1 == "First Jet" || fighterJet2 == "Second Jet" {
			err = render.RenderTemplate(w, r, "compare.page.tmpl", models.TemplateData{
				Error: "Choose valid aircrafts to compare",
				Data:  data,
			})
			if err != nil {
				log.Println(err)
			}
			return
		}

		jets, err := functions.GetTwoFighterJetsForCompare(fighterJet1, fighterJet2)
		if err != nil {
			log.Println(err)
		}

		postData := map[string]interface{}{}
		postData["first"] = jets[0]
		postData["second"] = jets[1]
		err = render.RenderTemplate(w, r, "compare.page.tmpl", models.TemplateData{
			PostData: postData,
			Data:     data,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
