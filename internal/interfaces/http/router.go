package http

import (
	"net/http"
	"todo-app/internal/interfaces/http/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(handler *handlers.TaskHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/tasks", handler.GetTasks)
	})

	r.Handle("/*", http.FileServer(http.Dir("./static/")))
	return r
}
