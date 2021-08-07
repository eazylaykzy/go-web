package handlers

import (
	"github.com/eazylaykzy/go-web/pkg/models"
	"github.com/eazylaykzy/go-web/pkg/render"
	"net/http"
)

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some business logics
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!!!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
