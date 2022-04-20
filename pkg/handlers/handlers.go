package handlers

import (
	"github.com/inuoluwadunsimi/bookings/pkg/config"
	"github.com/inuoluwadunsimi/bookings/pkg/models"
	"github.com/inuoluwadunsimi/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

//is the repository type
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}

}

// sets the repository for the handlers

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home-page handler

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIp", remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

// About is the about Page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Mr walton, come here I need to speak to you"

	remoteIp := m.App.Session.GetString(r.Context(), "remoteIp")
	stringMap["remoteIp"] = remoteIp

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
