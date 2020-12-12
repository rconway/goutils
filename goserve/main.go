package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
)

func main() {
	dir := "./"
	port := 8080

	h := http.FileServer(http.Dir(dir))
	h = handlers.CombinedLoggingHandler(os.Stdout, h)

	http.Handle("/", h)

	absDir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("ERROR: could not access directory - %v\n", dir)
	}

	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
	}()

	log.Printf("goserve: serving directory %v on port %v", absDir, port)

	c := make(chan int)
	<-c
}
