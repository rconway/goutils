package httputils

import (
	"fmt"
	"net/http"
	"strings"
)

// SubHandler type of a function that creates a Handler
// that optionally takes account of the provided path prefix.
type SubHandler func(prefix string) (handler http.Handler)

// SubHandlerFunc is a helper function to wrap a simple
// HandlerFunc as a SubHandler function.
func SubHandlerFunc(f http.HandlerFunc) SubHandler {
	return func(prefix string) http.Handler {
		return f
	}
}

// MuxSubGroup adds a sub-group '/$pattern/' to ServeMux taking account of path prefix
// for '/$pattern' -> '$prefix/$pattern/' redirections.
// prefix: path to prepend for an absolute path (needed for redirection)
// pattern: pattern to match in ServeMux
// handler: http.Handler for pattern
//
// Example usage...
// 	 // define a 'login' handler
// 	 func loginHandler(prefix string) http.Handler {
// 	 	mux := http.NewServeMux()
// 	 	httputils.MuxSubGroup(mux, prefix, "/", httputils.SubHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	 		fmt.Fprintf(w, "login handler\n")
// 	 	}))
// 	 	httputils.MuxSubGroup(mux, prefix, "/check", httputils.SubHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	 		fmt.Fprintf(w, "check handler\n")
// 	 	}))
// 	 	return mux
// 	 }
//
// 	 // define a handler for 'auth' that registers the login handler under subpath '/login'
// 	 func authHandler(prefix string) http.Handler {
// 	 	mux := http.NewServeMux()
// 	 	httputils.MuxSubGroup(mux, prefix, "/", httputils.SubHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	 		fmt.Fprintf(w, "auth handler\n")
// 	 	}))
// 	 	httputils.MuxSubGroup(mux, prefix, "/login", loginHandler)
// 	 	return mux
// 	 }
//
// 	 // define root handler that registers the auth handler under subpath '/auth' (which implies also '/auth/login')
// 	 root := http.NewServeMux()
// 	 httputils.MuxSubGroup(root, prefix, "/", httputils.SubHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	 	fmt.Fprintf(w, "root handler\n")
// 	 }))
// 	 httputils.MuxSubGroup(root, "", "/auth", authHandler)
//
func MuxSubGroup(mux *http.ServeMux, prefix string, pattern string, constructor SubHandler) {
	prefix = "/" + strings.Trim(prefix, "/")
	pattern = "/" + strings.Trim(pattern, "/")
	path := "/" + strings.Trim(prefix+pattern, "/")
	fmt.Printf("Mux Sub-Group: %v\n", path)
	handler := constructor(path)
	if pattern == "/" {
		mux.Handle(pattern, handler)
	} else {
		mux.Handle(pattern+"/", http.StripPrefix(pattern, handler))
		mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, path+"/", http.StatusMovedPermanently)
		})
	}
}
