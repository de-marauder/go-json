# go-json

## Introduction
This package expose a `MustParse` function which can be used to parse a JSON encoding to an appropriate go data structure.

It maps: 
- JSON strings to strings
- JSON numbers to numbers
- JSON booleans to booleans
- JSON arrays to slices
- JSON objects to maps
- JSON null to nil

## How to use
```bash
# Fetch package
go get github.com/de-marauder/go-json
go mod tidy
```

```go
// main.go

package main

import (
  "fmt"

  json "github.com/de-marauder/go-json"
) 

func main () {
  jsonStr := "[1, {\"name\": \"value\", \"arrKey\": [\"a\",3,\"w\"], \"objKey\": {\"nested key\": \"nested value\"} }, 3]"

  parsedJson := json.MustParse(jsonStr)

  fmt.Println(parsedJson)
}

```