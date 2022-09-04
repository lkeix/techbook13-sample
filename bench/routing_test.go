package bench

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/labstack/echo/v4"
	"github.com/lkeix/techbookfes13-sample/router"
)

var routes = []string{
	"/",
	"/hoge",
	"/fuga",
	"/health",
	"/hoge/fuga",
	"/n9te9",
	"/n9te9/hey",
	"/n9te9/hey/hello",
}

type value struct {
	str     string
	handler http.HandlerFunc
}

func BenchmarkEchoRouting(b *testing.B) {
	b.ReportAllocs()
	r := httptest.NewRequest("GET", "/", nil)
	u := r.URL
	w := httptest.NewRecorder()

	e := defineEchoRoute()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(routes); i++ {
			r.Method = http.MethodGet
			u.Path = routes[i]
			e.ServeHTTP(w, r)
		}
	}
}

func BenchmarkGinRouting(b *testing.B) {
	b.ReportAllocs()
	r := httptest.NewRequest("GET", "/", nil)
	u := r.URL
	w := httptest.NewRecorder()
	g := defineGinRoute()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(routes); i++ {
			r.Method = http.MethodGet
			u.Path = routes[i]
			g.ServeHTTP(w, r)
		}
	}
}

func BenchmarkChi(b *testing.B) {
	b.ReportAllocs()
	r := httptest.NewRequest("GET", "/", nil)
	u := r.URL
	w := httptest.NewRecorder()
	chir := defineChiRoute()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(routes); i++ {
			r.Method = http.MethodGet
			u.Path = routes[i]
			chir.ServeHTTP(w, r)
		}
	}
}

func BenchmarkMyRouter(b *testing.B) {
	b.ReportAllocs()
	r := httptest.NewRequest("GET", "/", nil)
	u := r.URL
	w := httptest.NewRecorder()
	myr := defineMyRoute()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(routes); i++ {
			r.Method = http.MethodGet
			u.Path = routes[i]
			myr.ServeHTTP(w, r)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

}

func defineMyRoute() *router.Router {
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
	}

	r := router.NewRouter()

	for _, v := range values {
		r.Insert(v.str, v.handler)
	}
	return r
}

func defineChiRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handler)
	r.Get("/hoge", handler)
	r.Get("/fuga", handler)
	r.Get("/health", handler)
	r.Get("/hoge/fuga", handler)
	r.Get("/{user}", handler)
	r.Get("/{user}/hey", handler)
	r.Get("/{user}/hey/{group}", handler)
	return r
}

func defineEchoRoute() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return nil
	})
	e.GET("/hoge", func(c echo.Context) error {
		return nil
	})
	e.GET("/fuga", func(c echo.Context) error {
		return nil
	})
	e.GET("/health", func(c echo.Context) error {
		return nil
	})
	e.GET("/hoge/fuga", func(c echo.Context) error {
		return nil
	})
	e.GET("/:user", func(c echo.Context) error {
		return nil
	})
	e.GET("/:user/hey", func(c echo.Context) error {
		return nil
	})
	e.GET("/:user/hey/:group", func(c echo.Context) error {
		return nil
	})
	return e
}

func defineGinRoute() *gin.Engine {
	g := gin.New()
	g.GET("/", func(ctx *gin.Context) {})
	g.GET("/hoge", func(ctx *gin.Context) {})
	g.GET("/fuga", func(ctx *gin.Context) {})
	g.GET("/health", func(ctx *gin.Context) {})
	g.GET("/hoge/fuga", func(ctx *gin.Context) {})
	g.GET("/:user", func(ctx *gin.Context) {})
	g.GET("/:user/hey", func(ctx *gin.Context) {})
	g.GET("/:user/hey/:group", func(ctx *gin.Context) {})
	return g
}
