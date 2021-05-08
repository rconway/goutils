package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/handlers"
)

var defaultPort = 8080

func getDefaultPort() (port int) {
	portStr, ok := os.LookupEnv("PORT")
	if ok {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			port = defaultPort
		}
	} else {
		port = defaultPort
	}
	return
}

func processCommandline(port *int) {
	flag.IntVar(port, "port", getDefaultPort(), "port to listen on")
	flag.Parse()
}

func main() {

	var port int
	processCommandline(&port)

	dir := "./"

	h := http.FileServer(http.Dir(dir))
	h = handlers.CombinedLoggingHandler(os.Stdout, h)

	http.Handle("/", h)

	absDir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("ERROR: could not access directory - %v\n", dir)
	}

	// start listening async
	go http.ListenAndServe(fmt.Sprint(":", port), nil)
	log.Printf("goserve: serving directory %v on port %v", absDir, port)

	// wait forever
	c := make(chan int)
	<-c
}
