package server

import (
	"strings"
	"testing"
)

func TestSanitizeAPODHTML(t *testing.T) {
	cases := []struct{ name, in, want string }{
		{"link kept", `<a href="https://x.org/a?b=1">hi</a>`, `<a href="https://x.org/a?b=1" target="_blank" rel="noopener noreferrer">hi</a>`},
		{"strong and br", `<strong>E:</strong> t<br>u`, `<strong>E:</strong> t<br>u`},
		{"script dropped", `a<script>alert(1)</script>b`, `ab`},
		{"js url unwrapped", `<a href="javascript:alert(1)">x</a>`, `x`},
		{"local path unwrapped", `<a href="///Users/me/a.html">Sunday</a>`, `Sunday`},
		{"onclick dropped", `<strong onclick="x()">s</strong>`, `<strong>s</strong>`},
		{"text escaped", `1 < 2 & 3`, `1 &lt; 2 &amp; 3`},
	}
	for _, c := range cases {
		if got := string(sanitizeAPODHTML(c.in)); got != c.want {
			t.Errorf("%s: got %q want %q", c.name, got, c.want)
		}
	}
	if strings.Contains(string(sanitizeAPODHTML(`<ahref="http://a.b">x</ahref>`)), "<a") {
		t.Error("malformed anchor should not produce a link")
	}
}
