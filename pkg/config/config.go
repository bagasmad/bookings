package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//doesn't import anything other than it absolutely has to
//danger import cycle not use any package of this app, only use standard package

// holds the application config, any site wide config should be put here
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
