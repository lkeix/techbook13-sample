package main

import (
	"net/http"

	"github.com/lkeix/techbookfes13-sample/router"
)

type value struct {
	str     string
	handler http.HandlerFunc
}

func main() {
	values := []value{
		{"/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("root"))
		}},
		{"/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("health"))
		}},
		{"/hey", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hey"))
		}},
		{"/:user/hey", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("health"))
		}},
		{"/:user", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok!"))
		}},
		{"/hoge", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hoge"))
		}},
		{"/fuga", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("fuga"))
		}},
		{"/hoge/fuga", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hoge fuga"))
		}},
	}

	r := router.NewRouter()

	for _, v := range values {
		r.Insert(v.str, v.handler)
	}

	http.ListenAndServe(":8080", r)
}
