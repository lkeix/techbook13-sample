package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lkeix/techbook13-sample/router"
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
		{"/hoge", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hoge"))
		}},
		{"/fuga", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("fuga"))
		}},
		{"/hoge/fuga", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hoge fuga"))
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
		{"/foo/:bar", func(w http.ResponseWriter, r *http.Request) {
			bar := router.Param(r, "bar")
			io.WriteString(w, bar)
		}},
		{"/foo/:bar/:baz/:qux/:quux/:corge", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("hogehoge")
			bar := router.Param(r, "bar")
			io.WriteString(w, bar)
			baz := router.Param(r, "baz")
			io.WriteString(w, baz)
			qux := router.Param(r, "qux")
			io.WriteString(w, qux)
			quux := router.Param(r, "quux")
			io.WriteString(w, quux)
			corge := router.Param(r, "corge")
			io.WriteString(w, corge)
		}},
		{"/foo/:bar/:baz/:qux/:quux/:corge/:grault/:garply/:waldo/:fred/:plugh", func(w http.ResponseWriter, r *http.Request) {
			bar := router.Param(r, "bar")
			io.WriteString(w, bar)
			baz := router.Param(r, "baz")
			io.WriteString(w, baz)
			qux := router.Param(r, "qux")
			io.WriteString(w, qux)
			quux := router.Param(r, "quux")
			io.WriteString(w, quux)
			corge := router.Param(r, "corge")
			io.WriteString(w, corge)
			grault := router.Param(r, "grault")
			io.WriteString(w, grault)
			garply := router.Param(r, "garply")
			io.WriteString(w, garply)
			waldo := router.Param(r, "waldo")
			io.WriteString(w, waldo)
			fred := router.Param(r, "fred")
			io.WriteString(w, fred)
			plugh := router.Param(r, "plugh")
			io.WriteString(w, plugh)
		}},
	}

	r := router.NewRouter()

	for _, v := range values {
		r.Insert(v.str, v.handler)
	}

	http.ListenAndServe(":8080", r)
}
