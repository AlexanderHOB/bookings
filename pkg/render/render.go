package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AlexanderHOB/bookings/pkg/config"
	"github.com/AlexanderHOB/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets a config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}
func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// Createa a template cache
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get requested template from myCache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couldn't not get template from cache")
	}
	buff := new(bytes.Buffer)
	td = addDefaultData(td)
	err := t.Execute(buff, td)
	if err != nil {
		log.Println(err)
	}
	// Render the template
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files named *.pages.html from templates
	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
