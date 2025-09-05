# httpez

httpez is a lightweight, user-friendly wrapper around Go’s standard net/http client. It makes managing headers and common request patterns easier while staying fully compatible with the standard library.

## Features
- Drop-in replacement for `http.Client`
- Global headers applied to all requests.
- Convenient `DoAndReadBody` helper to fetch response bodies.

## Installation

```bash
go get github.com/etaaa/httpez
```

## Usage
```golang
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
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests for bugs, feature requests, or improvements.