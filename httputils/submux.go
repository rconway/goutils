package httputils

import (
	"net/http"
	"strings"
)

// SubMux is a wrapper for http.ServeMux that supports a path prefix.
// This allows ServeMux instances to service subpaths.
// Example usage...
//   zzz
type SubMux struct {
	prefix string
	mux    *http.ServeMux
}

// NewSubMux constructs a SubMux with the supplied prefix.
func NewSubMux(prefix string) *SubMux {
	return NewSubWithMux(prefix, http.NewServeMux())
}

// NewSubWithMux constructs a SubMux with the supplied prefix,
// and existing ServceMux and re-using
func NewSubWithMux(prefix string, mux *http.ServeMux) *SubMux {
	return &SubMux{
		prefix: "/" + strings.Trim(prefix, "/"),
		mux:    mux,
	}
}

// Handle is a wrapper around http.ServeMux.Handle.
func (sub *SubMux) Handle(pattern string, handler http.Handler) {
	sub.mux.Handle(sub.prefix+"/"+strings.TrimLeft(pattern, "/"), handler)
}

// HandleFunc is a wrapper around http.ServeMux.HandleFunc.
func (sub *SubMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	sub.Handle(pattern, http.HandlerFunc(handler))
}

// Handler  is a wrapper around http.ServeMux.Handler.
func (sub *SubMux) Handler(r *http.Request) (h http.Handler, pattern string) {
	return sub.mux.Handler(r)
}

// ServeHTTP  is a wrapper around http.ServeMux.ServeHTTP.
func (sub *SubMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sub.mux.ServeHTTP(w, r)
}
