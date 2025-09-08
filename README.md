# httpez

[![Go Test](https://github.com/etaaa/httpez/actions/workflows/test.yml/badge.svg)](https://github.com/etaaa/httpez/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/etaaa/httpez)](https://goreportcard.com/report/github.com/etaaa/httpez)
[![Go Reference](https://pkg.go.dev/badge/github.com/etaaa/httpez.svg)](https://pkg.go.dev/github.com/etaaa/httpez)
[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

httpez is a lightweight, user-friendly wrapper around Go’s standard net/http client. It simplifies common HTTP tasks with a fluent, chainable API while remaining fully compatible with the standard library.

## Features

- **Fluent API**: Build clients and requests with a clean, chainable interface (WithBaseURL, WithHeader, WithQuery, etc.).

- **Default Configuration**: Easily set a BaseURL and default Headers that apply to all requests from a client.

- **Simple Payloads**: Effortlessly send JSON (WithJSON) or form data (WithForm) payloads. httpez handles the serialization and Content-Type headers for you.

- **Easy Response Handling**: Use helpers like AsBytes() and AsJSON() to quickly read response bodies without the boilerplate.

- **Extensible Middleware**: Wrap the client's transport with custom middleware for logging, authentication, retries, and more.

## Installation

```bash
go get github.com/etaaa/httpez
```

## Usage

This example shows how to create a client, set global headers, and perform a GET request with a query parameter.

```golang
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
```

For more detailed code samples covering a range of features, please see the `examples` folder.

## Contributing

Contributions are welcome! Please open issues or submit pull requests for bugs, feature requests, or improvements.
