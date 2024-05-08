package lodestone

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func has_class(node *html.Node, class string) bool {
	for _, attr := range node.Attr {
		if attr.Key == "class" {
			// Split the class attribute to handle multiple classes
			classes := strings.Split(attr.Val, " ")
			for _, c := range classes {
				if c == class {
					return true
				}
			}
		}
	}
	return false
}

func parse_html(body io.Reader, tag string, class string) string {
	doc, err := html.Parse(body)
	if err != nil {
		log.Fatal(err)
	}
	return traverse_nodes(doc, tag, class)
}

func traverse_nodes(node *html.Node, tagName string, className string) string {
	// Check if current node is the one we're looking for
	if node.Type == html.ElementNode && node.Data == tagName && has_class(node, className) {
		return extract_text(node) // Return the extracted text if it matches
	}

	// Continue searching in child nodes
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result := traverse_nodes(c, tagName, className) // Recurse into each child
		if result != "" {
			return result // Return only if we found a non-empty result
		}
	}

	return "" // Return empty if nothing is found throughout the subtree
}

func extract_text(node *html.Node) string {
	var text string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += strings.TrimSpace(c.Data) + " "
		}
	}
	return strings.TrimSpace(text)
}

func Get_FC_Name(id string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/freecompany/%s/", id))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		resp.Body.Close()
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	name := parse_html(resp.Body, "p", "entry__freecompany__name")

	return name, nil
}
