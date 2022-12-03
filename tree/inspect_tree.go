package tree

import (
	"fmt"
	"net/http"
)

type (
	kind        uint8
	routeMethod struct {
		ppath   string
		pnames  []string
		handler http.HandlerFunc
	}
)

const (
	staticKind kind = iota
	paramKind
	anyKind
)

func (n *Node) inspectAdd(method, path string, h http.HandlerFunc) {
	if path == "" {
		path = "/"
	}

	if path[0] != '/' {
		path = "/" + path
	}

	pnames := []string{}
	ppath := path

	if h == nil {
		panic("HandlerFunc is nil")
	}

	for i, lcpIndex := 0, len(path); i < lcpIndex; i++ {
		fmt.Printf("i: %d, lcpIndex: %d\n", i, lcpIndex)
		// is path param
		if path[i] == ':' {
			if i > 0 && path[i-1] == '\\' {
				path = path[:i-1] + path[i:]
				i--
				lcpIndex--
				continue
			}

			j := i + 1
			n.inspectInsert(method, path[:i], staticKind, routeMethod{})
			for ; i < lcpIndex && path[i] != '/'; i++ {
			}

			pnames = append(pnames, path[j:i])
			path = path[:j] + path[i:]
			i, lcpIndex = j, len(path)

			if i == lcpIndex {
				n.inspectInsert(method, path[:i], paramKind, routeMethod{ppath, pnames, h})
			} else {
				n.inspectInsert(method, path[:i], paramKind, routeMethod{})
			}
		} else if path[i] == '*' {
			n.inspectInsert(method, path[:i], staticKind, routeMethod{})
			pnames = append(pnames, "*")
			n.inspectInsert(method, path[:i+1], anyKind, routeMethod{ppath, pnames, h})
		}
	}

	n.inspectInsert(method, path, staticKind, routeMethod{ppath, pnames, h})
}

func (n *Node) inspectInsert(method, path string, t kind, rm routeMethod) {

}
