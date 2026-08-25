package colour

import "testing"

func TestPodColour(t *testing.T) {
	cases := map[string]string{
		"blue-7d9f9c5db4-abcde":  "blue",
		"green-6b8f7d4c9f-zzzzz": "green",
		"standalone":             "standalone",
		"":                       "",
	}
	for hostname, want := range cases {
		if got := PodColour(hostname); got != want {
			t.Errorf("PodColour(%q) = %q, want %q", hostname, got, want)
		}
	}
}

func TestCircle(t *testing.T) {
	if got := Circle("blue"); got != "🔵" {
		t.Errorf("Circle(%q) = %q, want 🔵", "blue", got)
	}
	if got := Circle("not-a-colour"); got != "" {
		t.Errorf("Circle(%q) = %q, want empty string", "not-a-colour", got)
	}
}

func TestNamespaceFromEnv(t *testing.T) {
	t.Setenv("NAMESPACE", "workshop")
	if got := Namespace(); got != "workshop" {
		t.Errorf("Namespace() = %q, want %q", got, "workshop")
	}
}

func TestNamespaceEmptyOutsideCluster(t *testing.T) {
	t.Setenv("NAMESPACE", "")
	if got := Namespace(); got != "" {
		t.Errorf("Namespace() = %q, want empty string outside a cluster", got)
	}
}
