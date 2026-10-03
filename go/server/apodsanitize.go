package server

import (
	"html/template"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var apodAllowedTags = map[string]bool{
	"a": true, "strong": true, "b": true, "em": true, "i": true, "br": true, "p": true,
}

// sanitizeAPODHTML keeps a small set of formatting tags and safe http(s) links,
// dropping everything else, so the result can be rendered as trusted HTML.
func sanitizeAPODHTML(markup string) template.HTML {
	context := &html.Node{Type: html.ElementNode, DataAtom: atom.Div, Data: "div"}
	fragments, err := html.ParseFragment(strings.NewReader(markup), context)
	if err != nil {
		return template.HTML(template.HTMLEscapeString(strings.TrimSpace(markup)))
	}

	var out strings.Builder
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			out.WriteString(template.HTMLEscapeString(n.Data))
			return
		case html.ElementNode:
			if !apodAllowedTags[n.Data] {
				if n.Data == "script" || n.Data == "style" {
					return
				}
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
				return
			}
			if n.Data == "br" {
				out.WriteString("<br>")
				return
			}
			open := "<" + n.Data + ">"
			if n.Data == "a" {
				href := safeAPODHref(n)
				if href == "" {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						walk(c)
					}
					return
				}
				open = `<a href="` + template.HTMLEscapeString(href) + `" target="_blank" rel="noopener noreferrer">`
			}
			out.WriteString(open)
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				walk(c)
			}
			out.WriteString("</" + n.Data + ">")
		}
	}
	for _, f := range fragments {
		walk(f)
	}
	return template.HTML(strings.TrimSpace(out.String()))
}

func safeAPODHref(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key != "href" {
			continue
		}
		u, err := url.Parse(strings.TrimSpace(a.Val))
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
			return ""
		}
		return u.String()
	}
	return ""
}
