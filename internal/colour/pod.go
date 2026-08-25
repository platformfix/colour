// Package colour serves the Kubernetes blue/green demo page: a background
// colour derived from the pod's hostname.
package colour

import (
	"os"
	"strings"
)

// circles maps a colour or namespace name to the emoji used to represent it
// on the served page. Unrecognised names have no circle.
var circles = map[string]string{
	"red":    "🔴",
	"orange": "🟠",
	"yellow": "🟡",
	"green":  "🟢",
	"blue":   "🔵",
	"purple": "🟣",
	"brown":  "🟤",
	"black":  "⚫",
	"white":  "⚪",
}

// Circle returns the emoji circle for name, or "" if name has none.
func Circle(name string) string {
	return circles[name]
}

// Hostname returns the pod's hostname, as set by Kubernetes via the
// HOSTNAME environment variable.
func Hostname() string {
	return os.Getenv("HOSTNAME")
}

// Namespace returns the current Kubernetes namespace: the NAMESPACE
// environment variable if set, otherwise the in-cluster serviceaccount
// namespace file. Returns "" outside a cluster with neither set.
func Namespace() string {
	if ns := os.Getenv("NAMESPACE"); ns != "" {
		return ns
	}
	data, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		return ""
	}
	return string(data)
}

// PodColour derives the demo colour from a hostname: the segment before the
// first "-", matching how Kubernetes names pods after their Deployment
// ("blue-7d9f9c5db4-abcde" -> "blue").
func PodColour(hostname string) string {
	return strings.SplitN(hostname, "-", 2)[0]
}
