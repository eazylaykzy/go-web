package handlers

import (
	"github.com/eazylaykzy/go-web/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.tmpl")
}
