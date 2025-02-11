package handlers

import (
	"log"
	"net/http"
)

type Handlers struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func (handlers *Handlers) Healthcheck(response http.ResponseWriter, request *http.Request) {
	handlers.logger.Println("A client hit the '/v1/healthcheck' endpoint")
	response.Write([]byte("Ok!"))
}
