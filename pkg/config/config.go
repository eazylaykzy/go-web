package config

import (
	"github.com/alexedwards/scs/v2"
	"text/template"
)

// AppConfig holds the application configurations
type AppConfig struct {
	UseCache      bool
	InProduction  bool
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
}
