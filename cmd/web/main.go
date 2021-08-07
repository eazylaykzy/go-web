package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/eazylaykzy/go-web/pkg/config"
	"github.com/eazylaykzy/go-web/pkg/handlers"
	"github.com/eazylaykzy/go-web/pkg/render"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

var (
	app     = config.AppConfig{}
	session *scs.SessionManager
)

func main() {
	// change this to true for production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	ts, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = ts
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println(fmt.Sprintf("Server started on port:%s", port))

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
