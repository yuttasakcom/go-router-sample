package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func usersRouter(s *mux.Router) {
	s.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api users"))
	})
}
