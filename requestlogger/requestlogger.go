package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("request logger")

	mux := mux.NewRouter()

	mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL:", r.URL)
		fmt.Fprintln(w, "Headers:")
		for k, v := range r.Header {
			fmt.Fprintf(w, "  %v: %v\n", k, v)
		}
	})

	log.Fatal(http.ListenAndServe(":80", mux))
}
