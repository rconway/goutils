package httputils

import (
	"net/http"
	"strings"
)

// SubMux zzz
type SubMux struct {
	prefix string
	mux *http.ServeMux
}

// NewSubMux zzz
func NewSubMux(prefix string) *SubMux {
	return &SubMux{
		prefix: "/" + strings.Trim(prefix, "/"),
		mux: http.NewServeMux(),
	}
}

// Handle zzz
func (sub *SubMux) Handle(pattern string, handler http.Handler) {
	sub.mux.Handle(sub.prefix + "/" + strings.TrimLeft(pattern, "/"), handler)
}

// HandleFunc zzz
func (sub *SubMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	sub.Handle(pattern, http.HandlerFunc(handler))
}

// Handler zzz
func (sub *SubMux) Handler(r *http.Request) (h http.Handler, pattern string) {
	return sub.mux.Handler(r)
}

// ServeHTTP zzz
func (sub *SubMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sub.mux.ServeHTTP(w, r)
}
