package colour

import (
	"html/template"
	"log/slog"
	"net/http"
)

var pageTemplate = template.Must(template.New("page").Parse(`<!DOCTYPE html>
<html>
<body style="background: {{.Namespace}}; text-align: center;">
<div style="padding: 4em;"></div>
<span style="padding: 4em; background: {{.PodColour}};">
<span style="padding: 2px; background: white;">
{{.Circles}}This is {{.DisplayName}}, serving {{.RequestPath}} for {{.RemoteAddr}}.
</span>
</span>
</body>
</html>
`))

type pageData struct {
	Namespace   string
	PodColour   string
	Circles     string
	DisplayName string
	RequestPath string
	RemoteAddr  string
}

// Handler serves the coloured demo page. Every field goes through
// html/template, which HTML-escapes it automatically — including the
// request-controlled RequestPath and RemoteAddr fields. jpetazzo/color built
// this same page with fmt.Sprintf against those two fields, unescaped; this
// is the fix for that reflected-HTML bug.
func Handler(w http.ResponseWriter, r *http.Request) {
	hostname := Hostname()
	namespace := Namespace()
	podColour := PodColour(hostname)

	displayName := hostname
	if namespace != "" {
		displayName = "pod " + namespace + "/" + hostname
	}

	data := pageData{
		Namespace:   namespace,
		PodColour:   podColour,
		Circles:     Circle(namespace) + Circle(podColour),
		DisplayName: displayName,
		RequestPath: r.URL.String(),
		RemoteAddr:  r.RemoteAddr,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pageTemplate.Execute(w, data); err != nil {
		slog.Error("render page", "error", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	slog.Info("request", "remote_addr", r.RemoteAddr, "method", r.Method, "path", r.URL.String())
}

// Healthz reports liveness/readiness for Kubernetes probes.
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
