package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func logger(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}

func main() {
	fmt.Println("request logger")

	mux := mux.NewRouter()
	mux.Use(logger)

	mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL:", r.URL)
		fmt.Fprintln(w, "Headers:")
		for k, v := range r.Header {
			fmt.Fprintf(w, "  %v: %v\n", k, v)
		}
	})

	log.Fatal(http.ListenAndServe(":80", mux))
}
