package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
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
		// RequestURI
		fmt.Fprintln(w, "RequestURI:", r.RequestURI)
		// Host
		fmt.Fprintln(w, "Host:", r.Host)
		// URL
		fmt.Fprintln(w, "URL:", r.URL)
		// Method
		fmt.Fprintln(w, "Method:", r.Method)
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
		// Body
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("ERROR reading request body:", err)
		} else {
			fmt.Fprintln(w, "Body:", string(data))
		}
	})
}

type summaryHandler struct{}

func (summaryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Host
	fmt.Fprintln(w, "Host:", r.Host)
	// URL
	fmt.Fprintln(w, "URL:", r.URL)
	// Method
	fmt.Fprintln(w, "Method:", r.Method)
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
	// Body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR reading request body:", err)
	} else {
		fmt.Fprintln(w, "Body:", string(data))
	}
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
