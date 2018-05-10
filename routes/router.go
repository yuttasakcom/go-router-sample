package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuttasakcom/gorilla-router-sample/middleware"
)

// Router http.Handler
func Router() http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.Logger)

	s := r.PathPrefix("/api").Subrouter()
	usersAPI(s)
	usersPostAPI(s)
	postsAPI(s)

	return r
}
