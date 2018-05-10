package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func usersPostAPI(s *mux.Router) {
	s.HandleFunc("/users/{userId}/post", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("POST /api/users/" + vars["userId"] + "/post"))
	}).Methods("POST")

	s.HandleFunc("/users/{userId}/posts", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("GET /api/users/" + vars["userId"] + "/posts"))
	}).Methods("GET")

	s.HandleFunc("/users/{userId}/post/{postId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("GET /api/users/" + vars["userId"] + "/post/" + vars["postId"]))
	}).Methods("GET")

	s.HandleFunc("/users/{userId}/post/{postId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("PUT /api/users/" + vars["userId"] + "/post/" + vars["postId"]))
	}).Methods("PUT")

	s.HandleFunc("/users/{userId}/post/{postId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("DELETE /api/users/" + vars["userId"] + "/post/" + vars["postId"]))
	}).Methods("DELETE")
}
