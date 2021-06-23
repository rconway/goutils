package httputils

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Masterminds/sprig"
)

var templateStr = `
{{- /* eat whitespace */ -}}
Host: {{ .R.Host }}
URL: {{ .R.URL }}
Method: {{ .R.Method }}
Headers:
{{- range $headerKey, $headerVal := .R.Header  }}
  {{ $headerKey }}: {{ $headerVal }}
{{- end }}
Params:
{{- range $paramKey, $paramVal := .R.URL.Query }}
  {{ $paramKey }}: {{- range $paramVal }} {{ . -}} {{ end }}
{{- end }}
Body:
{{ .Body | indent 2 }}
`

type requestDetails struct {
	R    *http.Request
	Body string
}

/*
DumpRequest provides a helper function to dump the received request
to the output writer, e.g dump to standard-out...

  httputils.DumpRequest(os.stdout, r)
*/
func DumpRequest(w io.Writer, r *http.Request) {
	tmpl, err := template.New("dump").Funcs(sprig.FuncMap()).Parse(templateStr)
	if err != nil {
		fmt.Fprintln(w, "ERROR parsing template:", err)
		return
	}

	// Prepare request details
	requestDetails := requestDetails{}
	requestDetails.R = r

	// Add the Body data
	bodyBytes, err := ioutil.ReadAll(r.Body)
	r.Body.Close() //  must close
	if err != nil {
		requestDetails.Body = fmt.Sprint("ERROR reading request body:", err)
	} else {
		requestDetails.Body = string(bodyBytes)
		// Reset the body data for a future user of this request
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	fmt.Fprintln(w, "zzz")
	err = tmpl.Execute(w, requestDetails)
	fmt.Fprintln(w, "zzz")
	if err != nil {
		fmt.Fprintln(w, "ERROR executing template:", err)
		return
	}
}

/*
DumpRequest provides a helper function to dump the received request
to the output writer, e.g dump to standard-out...

  httputils.DumpRequest(os.stdout, r)
*/
func xDumpRequest(w io.Writer, r *http.Request) {
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
