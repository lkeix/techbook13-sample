package tree

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
	}

	n := NewNode()

	for _, test := range tests {
		n.inspectAdd(test.method, test.path, test.h)
	}
}
