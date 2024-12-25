package render

import (
	"errors"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/srisudarshanrg/go-fighter-jets/server/config"
	"github.com/srisudarshanrg/go-fighter-jets/server/models"
)

var appConfig config.AppConfig

func AppConfigAccessRender(a config.AppConfig) {
	appConfig = a
}

// RenderTemplate executes a template from the template cache based on the passed template name
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData models.TemplateData) error {
	var templateCache map[string]*template.Template
	var err error

	if appConfig.UseTemplateCache {
		templateCache = appConfig.TemplateCache
	} else {
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}

	template, ok := templateCache[tmpl]
	if !ok {
		return errors.New("could not get template from template cache")
	}

	err = template.Execute(w, templateData)
	if err != nil {
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache with all the templates and layouts
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	files, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		templateName := filepath.Base(file)
		templateSet, err := template.New(templateName).ParseFiles()
		if err != nil {
			return templateCache, err
		}

		layoutMatches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}

		if len(layoutMatches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[templateName] = templateSet
	}

	return templateCache, nil
}
