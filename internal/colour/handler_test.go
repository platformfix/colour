package colour

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerEscapesRequestPath(t *testing.T) {
	t.Setenv("HOSTNAME", "blue-abc123")
	// The payload lives in the query string, not the path: net/url
	// percent-encodes "<"/">" in a URL path before Handler ever sees it, so a
	// path-based payload can't reach the template as literal HTML and would
	// pass here whether or not the handler escapes its output. RawQuery is
	// preserved verbatim in r.URL.String(), so this is the request-controlled
	// input that actually exercises the escaping.
	req := httptest.NewRequest(http.MethodGet, "/?<script>alert(1)</script>", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	body := rec.Body.String()
	if strings.Contains(body, "<script>") {
		t.Fatalf("response contains an unescaped <script> tag from the request path: %s", body)
	}
	if !strings.Contains(body, "&lt;script&gt;") {
		t.Fatalf("expected the request path to be HTML-escaped in the response, got: %s", body)
	}
}

func TestHandlerEscapesRemoteAddr(t *testing.T) {
	t.Setenv("HOSTNAME", "blue-abc123")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = `"><script>alert(1)</script>`
	rec := httptest.NewRecorder()

	Handler(rec, req)

	if strings.Contains(rec.Body.String(), "<script>") {
		t.Fatalf("response contains an unescaped <script> tag from RemoteAddr: %s", rec.Body.String())
	}
}

func TestHandlerReportsPodColour(t *testing.T) {
	t.Setenv("HOSTNAME", "blue-abc123")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	if !strings.Contains(rec.Body.String(), "blue-abc123") {
		t.Fatalf("expected the response to mention the hostname, got: %s", rec.Body.String())
	}
}

func TestHandlerDisplaysNamespaceAndPodColour(t *testing.T) {
	t.Setenv("HOSTNAME", "blue-7d9f9c5db4-abcde")
	t.Setenv("NAMESPACE", "green")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "pod green/blue-7d9f9c5db4-abcde") {
		t.Fatalf("expected response to contain the %q display format, got: %s", "pod <namespace>/<hostname>", body)
	}
	if !strings.Contains(body, Circle("green")) {
		t.Fatalf("expected response to contain the namespace's emoji circle %q, got: %s", Circle("green"), body)
	}
}

func TestHealthz(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	Healthz(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Healthz() status = %d, want %d", rec.Code, http.StatusOK)
	}
}
