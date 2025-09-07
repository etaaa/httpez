package main

import (
	"fmt"
	"log"

	"github.com/etaaa/httpez"
)

// Example struct to represent the JSON payload.
type CoffeeOrder struct {
	CoffeeType string `json:"coffeeType"`
	SugarCubes int    `json:"sugarCubes"`
}

func main() {
	// Create a new client.
	client := httpez.NewClient()

	// Prepare the JSON payload to be sent in the request body.
	payload := CoffeeOrder{
		CoffeeType: "Melange",
		SugarCubes: 1,
	}

	// Perform a POST request. The WithHeader and WithJSON methods are used
	// to fluently configure the request with a custom header and a JSON body.
	// WithJSON automatically handles marshaling the payload and setting the
	// correct Content-Type header.
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
