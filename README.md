# httpez

httpez is a lightweight, user-friendly wrapper around Go’s standard net/http client. It simplifies HTTP interactions while remaining fully compatible with the standard library.

## Features
- Drop-in replacement for `http.Client`
- Global headers applied to all requests.
- Fluent API for building clients and requests with a clean, chainable interface.
- Convenient response helpers: Use methods like `AsBytes` and `AsJSON` to easily fetch response bodies.
- Extensible middleware system for request and response logic.

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
```

For more detailed examples, including JSON POST requests, middleware, or response parsing, please see the `examples` folder.

## Contributing

Contributions are welcome! Please open issues or submit pull requests for bugs, feature requests, or improvements.