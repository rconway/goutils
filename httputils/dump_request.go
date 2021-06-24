package httputils

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Masterminds/sprig"
)

var txtTemplate = `
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

var tmpl *template.Template

func init() {
	var err error
	tmpl = template.New("dump").Funcs(sprig.FuncMap())
	tmpl, err = tmpl.Parse(txtTemplate)
	if err != nil {
		tmpl, _ = tmpl.Parse(fmt.Sprint("ERROR parsing template:", err))
		return
	}
}

type requestDetails struct {
	R    *http.Request
	Body string
}

func dumpRequestImpl(templateStr string, w io.Writer, r *http.Request) {
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
DumpRequestText provides a helper function to dump the received request
to the output writer in 'text' format, e.g dump to standard-out...

  httputils.DumpRequestText(os.stdout, r)
*/
func DumpRequestText(w io.Writer, r *http.Request) {
	dumpRequestImpl(txtTemplate, w, r)
}

/*
DumpRequest provides a helper function to dump the received request
to the output writer, e.g dump to standard-out...
Synonym for function DumpRequestText()

  httputils.DumpRequest(os.stdout, r)
*/
func DumpRequest(w io.Writer, r *http.Request) {
	DumpRequestText(w, r)
}
