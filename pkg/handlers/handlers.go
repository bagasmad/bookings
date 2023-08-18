package handlers

import (
	"net/http"

	"github.com/bagasmad/bookings/pkg/config"
	"github.com/bagasmad/bookings/pkg/models"
	"github.com/bagasmad/bookings/pkg/render"
)

// Repository used by the handlers
var Repo *Repository

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// SetRepo sets repository for the handlers
func SetRepo(r *Repository) {
	Repo = r
}

//create repository pattern, very common pattern that allows us to swap components of our app
//with minimal

type Repository struct {
	App *config.AppConfig
}

// Now every func that has the receiver will have access to the repository config
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//test session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["Home"] = "Hello, again this is home prett"
	stringMap["test"] = "Hello, again"

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
	//do business logic here, everytime the appropriate handler is called, we will return the appropriate template

	//we want to send data!!

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//pull IP address
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	//perform business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	stringMap["remote_ip"] = remoteIp
	//assign value to the template data, we are passing the stringMap
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
