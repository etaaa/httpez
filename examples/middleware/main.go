package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/etaaa/httpez"
)

func main() {
	// Create a new client and configure it with a custom middleware chain.
	// Middlewares are executed in the order they are added.
	client := httpez.NewClient().
		WithMiddleware(LoggerMiddleware()).
		WithMiddleware(AuthMiddleware("eyJhbGciOiJIUzI1NiIsInR5cCI6Ik..."))

	// Perform a GET request. This request will first pass through the
	// LoggerMiddleware, then the AuthMiddleware, before being sent.
	body, _, err := client.
		Get("https://httpbin.org/get").
		AsBytes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

// LoggerMiddleware logs the start and end of a request, including its duration.
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

// AuthMiddleware adds a bearer token to the Authorization header if it's not already set.
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
