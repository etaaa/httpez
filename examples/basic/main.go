package main

import (
	"fmt"
	"log"

	"github.com/etaaa/httpez"
)

func main() {
	// Create a new client and configure it with a base URL and default
	// headers. These settings will apply to all requests made with this
	// client.
	client := httpez.NewClient().
		WithBaseURL("https://httpbin.org").
		WithHeader("User-Agent", "httpez-example")

	// Performs a GET request using a relative path ("/get"). httpez
	// automatically combines this with the client's base URL to make a
	// request to "https://httpbin.org/get". The AsBytes() method reads
	// the entire response and closes the body.
	body, _, err := client.
		Get("/get").
		WithQuery("pastry", "apfelstrudel").
		AsBytes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
