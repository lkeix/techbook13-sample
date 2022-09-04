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
		{"/:user", func(w http.ResponseWriter, r *http.Request) {
			user := router.Param(r, "user")
			w.Write([]byte("ok! " + user))
		}},
		{"/:user/hey", func(w http.ResponseWriter, r *http.Request) {
			user := router.Param(r, "user")
			w.Write([]byte("hey " + user))
		}},
		{"/:user/hey/:group", func(w http.ResponseWriter, r *http.Request) {
			user := router.Param(r, "user")
			group := router.Param(r, "group")
			w.Write([]byte("hey " + user + " " + "group " + group))
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
