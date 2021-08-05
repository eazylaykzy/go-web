package handlers

import (
	"github.com/eazylaykzy/go-web/pkg/render"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.tmpl")
}
