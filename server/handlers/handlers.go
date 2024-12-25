package handlers

import (
	"log"
	"net/http"

	"github.com/srisudarshanrg/go-handlers-template/server/models"
	"github.com/srisudarshanrg/go-handlers-template/server/render"
)

// Home is the handler for the home page
func (app *HandlerRepository) Home(w http.ResponseWriter, r *http.Request) {
	err := render.RenderTemplate(w, "home.page.tmpl", models.TemplateData{})
	if err != nil {
		log.Println(err)
	}
}
