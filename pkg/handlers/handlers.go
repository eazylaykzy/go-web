package handlers

import "github.com/eazylaykzy/go-web/pkg/config"

// Repo is the repo used by the handlers
var Repo *Repository

// Repository is the repo type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handler
func NewHandlers(r *Repository) {
	Repo = r
}
