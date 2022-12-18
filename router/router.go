package router

import (
	"context"
	"net/http"
	"sync"

	"github.com/lkeix/techbook13-sample/tree"
)

type (
	Context struct {
		r       *http.Request
		w       http.ResponseWriter
		handler http.HandlerFunc
	}

	Router struct {
		root *tree.Node
		pool sync.Pool
	}
	ParamsKey string
)

const paramskey ParamsKey = "params"

func NewRouter() *Router {
	return &Router{
		root: tree.NewNode(),
		pool: sync.Pool{
			New: func() interface{} {
				return nil
			},
		},
	}
}

func (o *Router) Insert(path string, handler http.HandlerFunc) {
	o.root.Insert(path, handler)
}

func (o *Router) Search(path string) {
	o.root.Search(path)
}

func (o *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ctx.w = w
	// pathに対応するハンドラを探す
	handler, params := o.root.Search(r.URL.Path)

	if handler != nil {
		r = r.WithContext(context.WithValue(r.Context(), paramskey, params))
		handler(w, r)
		return
	}
}

func Param(r *http.Request, key string) string {
	params := r.Context().Value(paramskey).([]*tree.Param)
	for _, p := range params {
		if p.Key == key {
			return p.Value
		}
	}
	return ""
}
