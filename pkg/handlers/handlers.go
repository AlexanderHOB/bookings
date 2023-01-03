package handlers

import (
	"net/http"

	"github.com/AlexanderHOB/bookings/pkg/config"
	"github.com/AlexanderHOB/bookings/pkg/models"
	"github.com/AlexanderHOB/bookings/pkg/render"
)

//TemplateData hold data send from handlers to template
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the principal page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//About is all regarding of me
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	// perfomic some logic
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIp
	stringMap["test"] = "Hello from Huancayo"
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})

}
