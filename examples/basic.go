package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/etaaa/httpez"
)

func main() {
	client := httpez.NewClient()

	client.Headers().Set("User-Agent", "httpez-example")

	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		log.Fatal(err)
	}

	body, _, err := client.DoAndReadBody(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
