package httputils

import (
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

// SpaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The supplied filesystem is used to deliver
// the static content, which is assumed to be at the root of the filesystem.
type SpaHandler struct {
	Fsys fs.FS
}

// NewSpaHandler creates an SpaHandler instance whose static content is provided
// at the root of the supplied filesystem.
func NewSpaHandler(fsys fs.FS) *SpaHandler {
	return &SpaHandler{Fsys: fsys}
}

// ServeHTTP inspects the URL path to locate a file within the filesystem
// on the SPA handler. If a file is found, it will be served. If not, the path
// is adjusted to serve the root path, i.e. index.html. This is suitable behavior
// for serving an SPA (single page application).
func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Treat empty path as '/'
	if len(r.URL.Path) == 0 {
		r.URL.Path = "/"
	}

	// ensure we have a clean path
	r.URL.Path = path.Clean(r.URL.Path)

	// Trim leading '/', unless root path
	if len(r.URL.Path) > 1 {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	}

	// check whether a file exists at the given path
	// - excluding root path which is a synonym for the index file
	if r.URL.Path != "/" {
		_, err := fs.Stat(h.Fsys, r.URL.Path)
		if os.IsNotExist(err) {
			// If cannot stat file then use the root index file
			r.URL.Path = "/"
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// For root path, check that we have the index file.
	if r.URL.Path == "/" {
		_, err := fs.Stat(h.Fsys, "index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	// Use http.FileServer to serve the static dir
	http.FileServer(http.FS(h.Fsys)).ServeHTTP(w, r)
}
