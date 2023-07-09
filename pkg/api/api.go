package api

import (
	"log"
	"net/http"
	"thegoht/pkg/api/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartApi() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Route("/api", apiRouter)
    r.Get("/", routes.HTML)
    log.Fatal(http.ListenAndServe(":1234", r))
}

func apiRouter(r chi.Router) {
    r.Get("/", routes.Root)
    r.Route("/todo", todosRouter)
}

func todosRouter(r chi.Router) {
    r.Post("/", routes.AddTodo)
    r.Put("/", routes.EditTodo)
}
