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
		// URL
		fmt.Fprintln(w, "URL:", r.URL)
		// Headers
		fmt.Fprintln(w, "Headers:")
		for headerKey, headerVal := range r.Header {
			fmt.Fprintf(w, "  %v: %v\n", headerKey, headerVal)
		}
		// Query params...
		fmt.Fprintln(w, "Params:")
		for paramKey, paramVal := range r.URL.Query() {
			fmt.Fprint(w, "  ", paramKey, ":")
			for _, paramValItem := range paramVal {
				fmt.Fprint(w, " ", paramValItem)
			}
			fmt.Fprintln(w)
		}
	})

	log.Fatal(http.ListenAndServe(":80", mux))
}
