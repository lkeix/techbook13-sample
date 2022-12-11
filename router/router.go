package router

import (
	"context"
	"net/http"
	"sync"

	"github.com/lkeix/techbookfes13-sample/tree"
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

	routeMethod struct {
		ppath   string
		pnames  []string
		handler http.HandlerFunc
	}

	kind      uint8
	ParamsKey string
)

const (
	paramskey ParamsKey = "params"

	staticKind kind = iota
	paramKind
	anyKind
)

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

func (o *Router) InspectAdd(method, path string, handler http.HandlerFunc) {
	pnames := []string{}
	ppath := path

	if handler == nil {
		panic("handler is nil")
	}

	for i, lcpIndex := 0, len(ppath); i < lcpIndex; i++ {
		if path[i] == ':' {
			if i > 0 && path[i-1] == '\\' {
				path = path[:i-1] + path[i:]
				i--
				lcpIndex--
				continue
			}

			j := i + 1
			o.Inspectinsert(method, path[:i], staticKind, routeMethod{})
			for ; i < lcpIndex && path[i] != '/'; i++ {
			}

			pnames = append(pnames, path[j:i])
			path = path[:j] + path[i:]
			i, lcpIndex = j, len(path)

			if i == lcpIndex {
				o.Inspectinsert(method, path[:i], paramKind, routeMethod{ppath, pnames, handler})
			} else {
				o.Inspectinsert(method, path[:i], paramKind, routeMethod{})
			}
		} else if path[i] == '*' {
			o.Inspectinsert(method, path[:i], staticKind, routeMethod{})
			pnames = append(pnames, "*")
			o.Inspectinsert(method, path[:i+1], anyKind, routeMethod{ppath, pnames, handler})
		}
	}

	o.Inspectinsert(method, path, staticKind, routeMethod{ppath, pnames, handler})
}

func (o *Router) Inspectinsert(method, path string, t kind, rm routeMethod) {
	currentNode := o.root

	if currentNode == nil {
		panic("root is nil")
	}

	search := path
	for {
		searchLen := len(search)
		prefixLen := len(currentNode.Prefix)

		mx := prefixLen
		if searchLen < mx {
			mx = searchLen
		}

		lcpIndex := 0
		for ; lcpIndex < mx && search[lcpIndex] == currentNode.Prefix[lcpIndex]; lcpIndex++ {

		}

		if lcpIndex == 0 {

		} else if lcpIndex < prefixLen {

		} else if lcpIndex < searchLen {

		} else {

		}

		return
	}
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
