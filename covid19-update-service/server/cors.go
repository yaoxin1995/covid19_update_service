package server

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/gorilla/handlers"
)

var allowedHeaders = []string{"Accept", "Content-Type", "Content-Length"}

type corsHandler struct {
	allowedMethods []string
	wrapper        http.Handler
}

func cors(r *mux.Router) func(handlerFunc http.Handler) http.Handler {
	allowedMethods := getAllMethodsForRoute(r)
	return func(h http.Handler) http.Handler {
		return &corsHandler{
			allowedMethods: getAllMethodsForRoute(r),
			wrapper: handlers.CORS(handlers.AllowedHeaders(allowedHeaders),
				handlers.AllowedMethods(allowedMethods))(h),
		}
	}
}

func (c *corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Allowed", strings.Join(c.allowedMethods, ", "))
	}
	c.wrapper.ServeHTTP(w, r)
}
