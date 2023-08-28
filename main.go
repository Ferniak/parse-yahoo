package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://finance.yahoo.com/news/boeing-shares-fall-737-delivery-095832970.html"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка запроса: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	traverseDiv(doc)

}

func traverseDiv(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "time" || n.Data == "div") {
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == "caas-body" {
				divText := getTextContent(n)
				fmt.Println(divText)

			}
			if attr.Key == "class" && attr.Val == "caas-attr-meta-time" {
				divText := getTextContent(n)
				fmt.Println(divText)
				fmt.Println("Hello")
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseDiv(c)
	}
	fmt.Println("Hello")

}

func getTextContent(n *html.Node) string {
	var textContent string

	if n.Type == html.TextNode {
		textContent = n.Data
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			textContent += getTextContent(c)
		}
	}

	return strings.TrimSpace(textContent)
}
