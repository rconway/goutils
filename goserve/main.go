package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// router := mux.NewRouter()
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	// http.ListenAndServe(":8080", router)

	dir := "./"
	port := 8080

	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Printf("goserve: serving directory %v on port %v", dir, port)

	http.ListenAndServe(fmt.Sprint(":", port), nil)
}
