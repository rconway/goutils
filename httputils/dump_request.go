package httputils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

/*
DumpRequest provides a helper function to dump the received request
to the output writer, e.g dump to standard-out...

  httputils.DumpRequest(os.stdout, r)
*/
func DumpRequest(w io.Writer, r *http.Request) {
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
