package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func usersAPI(s *mux.Router) {
	s.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET /api/users"))
	}).Methods("GET")

	s.HandleFunc("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("GET /api/users/" + vars["userId"]))
	}).Methods("GET")

	s.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST /api/users"))
	}).Methods("POST")

	s.HandleFunc("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("PUT /api/users/" + vars["userId"]))
	}).Methods("PUT")

	s.HandleFunc("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("DELETE /api/users/" + vars["userId"]))
	}).Methods("DELETE")
}
