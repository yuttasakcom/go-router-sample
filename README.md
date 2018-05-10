# Gorilla Router Sample

## Table of Contents

* [maing](#main)
* [Routes](#routes)
* [router](#router)
* [usersAPI](#usersapi)
* [usersPostAPI](#userspostapi)
* [postsAPI](#postsapi)

## Main

```go
package main

import (
	"log"
	"net/http"

	"github.com/yuttasakcom/gorilla-router-sample/routes"
)

func main() {
	r := routes.Router()

	log.Fatal(http.ListenAndServe(":3000", r))
}
```

## Routes

## router

```go
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router http.Handler
func Router() http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/api").Subrouter()
	usersAPI(s)
	usersPostAPI(s)
	postsAPI(s)

	return r
}
```

## usersAPI

```go
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
```

## usersPostAPI

```go
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
```

## postsAPI

```go
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
```
