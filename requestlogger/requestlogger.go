package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/rconway/goutils/httputils"
)

var defaultPort uint16 = 80

func logger(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}

func getPort() (port uint16) {
	portEnv, ok := os.LookupEnv("PORT")
	if ok {
		portVal, err := strconv.ParseUint(portEnv, 10, 16)
		if err != nil {
			port = defaultPort
			log.Println("WARNING: Bad PORT specified. Using default port =", port)
		} else {
			port = uint16(portVal)
			log.Println("Using specified port =", port)
		}
	} else {
		port = defaultPort
		log.Println("Using default port =", port)
	}
	return port
}

func myHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httputils.DumpRequest(w, r)
	})
}

type summaryHandler struct{}

func (summaryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httputils.DumpRequest(w, r)
}

func main() {
	PORT := getPort()
	log.Println("request logger running on port:", PORT)

	// Create handler chain
	var handler http.Handler
	handler = myHandler()
	handler = logger(handler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), handler))
}
