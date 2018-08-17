package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// processElement will be called for each HTML element found
func processElement(_ int, element *goquery.Selection) {
	fmt.Println(element.Text())
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func main() {
	// Make HTTP request
	var query string
	fmt.Print("What do you want to search in GOOGLE: ")
	fmt.Scan(&query)
	escapedQuery := url.QueryEscape(query)
	response, err := http.Get("http://google.com/search?q=" + escapedQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	// Find all links and process them with the function
	// defined earlier
	document.Find("h3.r a:not(.l)").Each(processElement)
}
