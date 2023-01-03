package main

import (
	"net/http"

	"github.com/AlexanderHOB/bookings/pkg/config"
	"github.com/AlexanderHOB/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// Establecemos un file server para servir archivos estáticos desde la carpeta "./static/".
	fileServer := http.FileServer(http.Dir("../../static/"))
	// Asignamos el file server al enrutador de la aplicación, eliminando el prefijo "/static" de la ruta de la solicitud.
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
