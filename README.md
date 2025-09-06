# httpez

httpez is a lightweight, user-friendly wrapper around Go’s standard net/http client. It simplifies HTTP interactions while remaining fully compatible with the standard library.

## Features
- Drop-in replacement for `http.Client`
- Global headers applied to all requests.
- Extensible middleware system for request and response logic.
- Convenient `*AndReadBody` helpers for quickly fetching response bodies.
- Fluent API for building clients and requests

## Installation

```bash
go get github.com/etaaa/httpez
```

## Usage

Here is a basic example demonstrating how to create a client and make a simple request with headers.

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

	body, _, err := client.GetAndReadBody("https://httpbin.org/headers")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
```
For more detailed examples, including how to use middleware, please see the `examples` folder.

## Contributing

Contributions are welcome! Please open issues or submit pull requests for bugs, feature requests, or improvements.