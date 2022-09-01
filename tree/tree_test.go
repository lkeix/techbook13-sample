package tree

import (
	"net/http"
	"reflect"
	"testing"
)

type value struct {
	str     string
	handler http.HandlerFunc
}

func TestTree(t *testing.T) {
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

	root := NewNode()

	for _, v := range values {
		root.Insert(v.str, v.handler)
	}

	for _, v := range values {
		handler := root.Search(v.str)
		// 関数ポインタで比較する
		if reflect.ValueOf(handler).Pointer() != reflect.ValueOf(v.handler).Pointer() {
			t.Errorf("Expected: %v\n Actually: %v", v.handler, handler)
		}
	}
}
