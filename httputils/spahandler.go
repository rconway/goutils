package httputils

import (
	"net/http"
	"os"
	"path/filepath"
)

// SpaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type SpaHandler struct {
	Root       http.FileSystem
	PathPrefix string
	IndexPath  string
	StatFunc   func(name string) (os.FileInfo, error)
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.PathPrefix, path)

	// check whether a file exists at the given path
	_, err = h.StatFunc(path)
	if err != nil {
		// If cannot stat file then use the root index file
		r.URL.Path = filepath.Join(h.PathPrefix, h.IndexPath)
	}

	// Use http.FileServer to serve the static dir
	http.FileServer(h.Root).ServeHTTP(w, r)
}
