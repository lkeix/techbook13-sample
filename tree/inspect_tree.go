package tree

import (
	"bytes"
	"net/http"
)

type (
	Kind        uint8
	RouteMethod struct {
		ppath   string
		pnames  []string
		handler http.HandlerFunc
	}

	Children []*InspectNode

	InspectNode struct {
		Label          byte
		Prefix         string
		Kind           Kind
		Parent         *InspectNode
		StaticChildren Children
		OriginalPath   string
		Methods        *RouteMethods
		ParamChild     *InspectNode
		AnyChild       *InspectNode
		ParamsCount    int
		isLeaf         bool
		isHandler      bool
	}
)

type RouteMethods struct {
	Connect     *RouteMethod
	Delete      *RouteMethod
	Get         *RouteMethod
	Head        *RouteMethod
	Options     *RouteMethod
	Patch       *RouteMethod
	Post        *RouteMethod
	Propfind    *RouteMethod
	Put         *RouteMethod
	Trace       *RouteMethod
	Report      *RouteMethod
	allowHeader string
}

func (r *RouteMethods) updateAllowHandler() {
	buf := new(bytes.Buffer)
	buf.WriteString(http.MethodOptions)

	if r.Connect != nil {
		buf.WriteString(", ")
		buf.WriteString(http.MethodConnect)
	}

	if r.Delete != nil {
		buf.WriteString(", ")
		buf.WriteString(http.MethodDelete)
	}

	if r.Get != nil {
		buf.WriteString(", ")
		buf.WriteString(http.MethodGet)
	}

	if r.Head != nil {
		buf.WriteString(", ")
		buf.WriteString(http.MethodHead)
	}

	if r.Patch != nil {
		buf.WriteString(", ")
		buf.WriteString(http.MethodPatch)
	}

	if r.Post != nil {
		buf.WriteString(", ")
		buf.WriteString(http.MethodPost)
	}

	if r.Put != nil {
		buf.WriteString(", ")
		buf.WriteString(http.MethodPut)
	}

	if r.Trace != nil {
		buf.WriteString(", ")
		buf.WriteString(http.MethodTrace)
	}

	if r.Report != nil {
		buf.WriteString(", ")
		buf.WriteString("REPORT")
	}

	r.allowHeader = buf.String()
}

const (
	staticKind Kind = iota
	paramKind
	anyKind
)

func (n *InspectNode) AddMethod(method string, h *RouteMethod) {
	switch method {
	case http.MethodConnect:
		n.Methods.Connect = h
	case http.MethodDelete:
		n.Methods.Delete = h
	case http.MethodGet:
		n.Methods.Get = h
	case http.MethodHead:
		n.Methods.Head = h
	case http.MethodOptions:
		n.Methods.Options = h
	case http.MethodPatch:
		n.Methods.Patch = h
	case http.MethodPut:
		n.Methods.Put = h
	case http.MethodPost:
		n.Methods.Post = h
	case http.MethodTrace:
		n.Methods.Trace = h
	}
	n.Methods.updateAllowHandler()
	n.isHandler = true
}

func (n *InspectNode) InspectAdd(method, path string, h http.HandlerFunc) {
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
		// is path param
		if path[i] == ':' {
			if i > 0 && path[i-1] == '\\' {
				path = path[:i-1] + path[i:]
				i--
				lcpIndex--
				continue
			}

			j := i + 1
			n.inspectInsert(method, path[:i], staticKind, RouteMethod{})
			for ; i < lcpIndex && path[i] != '/'; i++ {
			}

			pnames = append(pnames, path[j:i])
			path = path[:j] + path[i:]
			i, lcpIndex = j, len(path)

			if i == lcpIndex {
				n.inspectInsert(method, path[:i], paramKind, RouteMethod{ppath, pnames, h})
			} else {
				n.inspectInsert(method, path[:i], paramKind, RouteMethod{})
			}
		} else if path[i] == '*' {
			n.inspectInsert(method, path[:i], staticKind, RouteMethod{})
			pnames = append(pnames, "*")
			n.inspectInsert(method, path[:i+1], anyKind, RouteMethod{ppath, pnames, h})
		}
	}

	n.inspectInsert(method, path, staticKind, RouteMethod{ppath, pnames, h})
}

func (n *InspectNode) inspectInsert(method, path string, t Kind, rm RouteMethod) {

}
