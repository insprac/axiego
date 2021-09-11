# Axiego

An experimental implementation of the Axie Marketplace GraphQL API in Go.

Note: this package currently only implements a few types and is not complete.

## Installation

To install the Axiego package you need to isntall Go and set your Go workspace
first.

1. You can use the below command to install Axiego.
```
$ go get -u github.com/insprac/axiego
```

2. Import it in your code:
```go
import "github.com/insprac/axiego"
```

## Quick Start

```go
package main

import (
	"log"

	"github.com/insprac/axiego"
)

const query = `
query GetLatestListings($from: Int, $size: Int) {
  axies(from: $from, size: $size, auctionType: Sale, sort: Latest) {
    total
    results {
      id
      name
      image
      class
      parts {
        id
        name
        type
        class
      }
      auction {
        currentPrice
        currentPriceUSD
      }
    }
  }
}
`

type Response struct {
	Data   ResponseData   `json:"data"`
	Errors []axiego.Error `json:"errors"`
}

type ResponseData struct {
	Axies axiego.Axies `json:"axies"`
}

func main() {
	vars := map[string]interface{}{
		"from": 0,
		"size": 10,
	}
	res := Response{}

	err := axiego.Req(query, vars, &res)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
```
