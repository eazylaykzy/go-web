package handlers

import (
	"github.com/eazylaykzy/go-web/pkg/models"
	"github.com/eazylaykzy/go-web/pkg/render"
	"net/http"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}
