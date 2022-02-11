package handler

import (
	"WSServer/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

// NewHandler - return a new http server
func NewHandler() http.Handler {
	sm := server.NewServerManager()

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	// r.Use(cors.AllowAll().Handler)

	r.HandleFunc("/connect", sm.SocketConnection)

	return r
}
