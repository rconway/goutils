package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

func main() {

	PORT := getPort()
	log.Println("request logger running on port:", PORT)

	mux := mux.NewRouter()
	mux.Use(logger)

	mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Host
		fmt.Fprintln(w, "Host:", r.Host)
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

	http.Handle("/", mux)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}
