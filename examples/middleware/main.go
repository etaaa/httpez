package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/etaaa/httpez"
)

func main() {
	// Create a new client.
	client := httpez.NewClient()

	// Use the custom middlewares. The order matters: LoggerMiddleware will
	// execute first, then AuthMiddleware, and finally the request will be sent.
	client.
		Use(LoggerMiddleware()).
		Use(AuthMiddleware("eyJhbGciOiJIUzI1NiIsInR5cCI6Ik..."))

	// Performs a GET request to the specified URL, reads and returns
	// the entire response body, and automatically closes the response body.
	body, _, err := client.GetAndReadBody("https://httpbin.org/headers")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func LoggerMiddleware() httpez.Middleware {
	return func(next http.RoundTripper) http.RoundTripper {
		return httpez.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			start := time.Now()
			log.Printf("Starting request to %s", req.URL.String())
			resp, err := next.RoundTrip(req)
			log.Printf("Finished request to %s in %v", req.URL.String(), time.Since(start))
			return resp, err
		})
	}
}

func AuthMiddleware(token string) httpez.Middleware {
	return func(next http.RoundTripper) http.RoundTripper {
		return httpez.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Authorization") == "" {
				req.Header.Set("Authorization", "Bearer "+token)
			}
			return next.RoundTrip(req)
		})
	}
}
