package main

import (
	"fmt"
	"log"

	"github.com/etaaa/httpez"
)

func main() {
	// Create a new client.
	client := httpez.NewClient()

	// Set custom headers for all requests made with this client.
	client.Headers().
		Set("Accept", "application/json").
		Set("User-Agent", "httpez-example")

	// Performs a GET request to the specified URL with a query parameter,
	// reads and returns the entire response body, and automatically closes
	// the response body.
	body, _, err := client.
		Get("https://httpbin.org/get").
		WithQuery("foo", "bar").
		AsBytes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
