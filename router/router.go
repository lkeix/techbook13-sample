package router

import (
	"net/http"

	"github.com/lkeix/techbookfes13-sample/tree"
)

type Router struct {
	root *tree.Node
}

func NewRouter() *Router {
	return &Router{
		root: tree.NewNode(),
	}
}

func (o *Router) Insert(path string, handler http.HandlerFunc) {
	o.root.Insert(path, handler)
}

func (o *Router) Search(path string) {
	o.root.Search(path)
}

func (o *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// pathに対応するハンドラを探す
	handler := o.root.Search(path)

	if handler != nil {
		handler(w, r)
		return
	}
}
