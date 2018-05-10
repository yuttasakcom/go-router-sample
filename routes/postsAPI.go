package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func postsAPI(s *mux.Router) {
	s.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET /api/posts"))
	}).Methods("GET")

	s.HandleFunc("/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("GET /api/posts/" + vars["postId"]))
	}).Methods("GET")
}
