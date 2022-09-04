package router

import (
	"context"
	"net/http"

	"github.com/lkeix/techbookfes13-sample/tree"
)

type (
	Router struct {
		root *tree.Node
	}
	ParamsKey string
)

const paramskey ParamsKey = "params"

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
	handler, params := o.root.Search(path)

	if handler != nil {
		r = r.WithContext(context.WithValue(r.Context(), paramskey, params))
		handler(w, r)
		return
	}
}

func Param(r *http.Request, key string) string {
	params := r.Context().Value(paramskey).([]tree.Param)
	for _, p := range params {
		if p.Key == key {
			return p.Value
		}
	}
	return ""
}
