package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline3(nil, doc)

}

//Выводит содержание html элемента
func outline3(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Println(n.Data, "current dom element")
		text := &bytes.Buffer{}
		collectText(n, text)
		fmt.Println(text)
		stack = append(stack, n.Data) // Внесение дескриптора в стек

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//Рекурсия
		outline3(stack, c)
	}
}

func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}
