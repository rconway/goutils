package httputils

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSubPath(t *testing.T) {
	t.Log("zzz")

	mux := http.NewServeMux()

	sm := NewSubWithMux("/user/", mux)
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this is /user")
	})
	sm.HandleFunc("/fred/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this is /user/fred")
	})

	http.ListenAndServe(":8080", mux)
}
