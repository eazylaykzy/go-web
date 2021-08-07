package render

import (
	"bytes"
	"fmt"
	"github.com/eazylaykzy/go-web/pkg/config"
	"github.com/eazylaykzy/go-web/pkg/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	functions = template.FuncMap{}
	app       *config.AppConfig
)

// AddDefaultData add default template data
func AddDefaultData(a *models.TemplateData) *models.TemplateData {
	return a
}

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func Template(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// get the template cache from the app config
	tc := map[string]*template.Template{}

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from temmplate cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, err
}
