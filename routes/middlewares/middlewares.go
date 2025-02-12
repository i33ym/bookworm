package middlewares

import (
	"log"
	"net/http"
)

type Middlewares struct {
	logger *log.Logger
}

func NewMiddlewares(logger *log.Logger) *Middlewares {
	return &Middlewares{
		logger: logger,
	}
}

func (middlewares *Middlewares) LogRequest(next http.Handler) http.Handler {
	fn := func (response http.ResponseWriter, request *http.Request)  {
		middlewares.logger.Println(request.Method, request.URL.Path)
		next.ServeHTTP(response, request)
	}

	return http.HandlerFunc(fn)
}