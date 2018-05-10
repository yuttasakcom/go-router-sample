package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router http.Handler
func Router() http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/api").Subrouter()
	usersRouter(s)

	return r
}
