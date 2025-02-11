package routes

import (
	"log"
	"net/http"

	"api.bookworm.cc/routes/handlers"
	"github.com/julienschmidt/httprouter"
)

type Routes struct {
	logger *log.Logger
}

func NewRoutes(logger *log.Logger) *Routes {
	return &Routes{
		logger: logger,
	}
}

func (routes *Routes) API() http.Handler {
	mux := httprouter.New()
	handlers := handlers.NewHandlers(routes.logger)

	mux.HandlerFunc(http.MethodGet, "/v1/healthcheck", handlers.Healthcheck)
	mux.HandlerFunc(http.MethodGet, "/v1/books/:id", handlers.ViewBook)
	mux.HandlerFunc(http.MethodPost, "/v1/books", handlers.CreateBook)

	return mux
}
