package router

import (
	"net/http"
	"testing"
)

func testHandler(w http.ResponseWriter, r *http.Request) {}

func TestInspectAdd(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		h      http.HandlerFunc
	}{
		{
			name:   "insert root static path",
			method: http.MethodGet,
			path:   "/",
			h:      testHandler,
		},
		{
			name:   "insert 1 nested static path",
			method: http.MethodGet,
			path:   "/hoge",
			h:      testHandler,
		},
		{
			name:   "insert 2 nested static path",
			method: http.MethodGet,
			path:   "/hoge/fuga",
			h:      testHandler,
		},
		{
			name:   "insert 1 nested path param path",
			method: http.MethodGet,
			path:   "/:user",
			h:      testHandler,
		},
		{
			name:   "insert 1 nested path param path, 1 nested static path",
			method: http.MethodGet,
			path:   "/:user/hoge",
			h:      testHandler,
		},
		{
			name:   "insert 1 nested static path, 1 nested param path",
			method: http.MethodGet,
			path:   "/hoge/:user",
			h:      testHandler,
		},
		{
			name:   "insert 1 nested any path",
			method: http.MethodGet,
			path:   "/hoge-*",
			h:      testHandler,
		},
		{
			name:   "insert 1 nested any path, 1 nested param path",
			method: http.MethodGet,
			path:   "/hoge-*/:user",
			h:      testHandler,
		},
	}

	r := NewRouter()

	for _, test := range tests {
		r.InspectAdd(test.method, test.path, test.h)
	}
}
