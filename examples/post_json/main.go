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
		Set("User-Agent", "httpez-example")

	// Prepare the JSON payload.
	payload := Payload{
		Name:  "httpez",
		Value: 1,
	}

	// Perform a POST request to the specified URL with the JSON payload,
	// read and return the entire response body, and automatically close
	// the response body.
	body, _, err := client.
		Post("https://httpbin.org/post", nil).
		WithHeader("X-Request-ID", "917dfcee-7155-416d-ab65-35b9e1a1ecd1").
		WithJSON(payload).
		AsBytes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

// Example struct to represent the JSON payload.
type Payload struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
