package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuttasakcom/gorilla-router-sample/middleware"
)

func usersPostAPI(s *mux.Router) {

	s.Handle("/users/{userId}/post", middleware.Chain(
		middleware.Auth("token"),
	)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("POST /api/users/" + vars["userId"] + "/posts"))
	}))).Methods("POST")

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

func userCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created"))
}
