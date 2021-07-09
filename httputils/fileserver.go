package httputils

import (
	"net/http"
	"path"
)

// DefaultFileContentFunc is a function type an implementation of which provides the default
// content from the given file-system to the given http response writer.
type DefaultFileContentFunc func(httpFs http.FileSystem, w http.ResponseWriter, r *http.Request)

// FileServerWithDefault returns an http handler that serves the requested URL path from the
// given file-system. In the case that the requested URL path is not found within the given
// file-system, then the default content is provided via the given function.
func FileServerWithDefault(httpFs http.FileSystem, defaultContent DefaultFileContentFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the file exists by attempting to open it
		file, err := httpFs.Open(path.Clean(r.URL.Path))
		if err != nil {
			// If can't open file then initiate the default action
			defaultContent(httpFs, w, r)
			return
		} else {
			// Else the file exists so serve it using a FileServer to deliver the content
			file.Close() // we don't need this
			http.FileServer(httpFs).ServeHTTP(w, r)
		}
	})
}
