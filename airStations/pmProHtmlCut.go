package airStations

import (
	"bytes"
	"errors"
	"io"

	"golang.org/x/net/html"
)

func body(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "div" && node.Attr != nil && node.Attr[0].Key == "id" && node.Attr[0].Val == "summary" {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}

	return nil, errors.New("Missing <body> in the node tree")
}

func renderNode(n *html.Node) string {
	//func renderNode(n *html.Node) []byte {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	//return buf.Bytes()
	html.Render(w, n)
	return buf.String()
}

func GetErrorMessageFromHtmlView(pmProResponse []byte) (errorFound string, err error) {
	doc, _ := html.Parse(bytes.NewBuffer(pmProResponse))
	bn, err := body(doc)
	if err != nil {
		return
	}
	errorFound = renderNode(bn)
	return
}
