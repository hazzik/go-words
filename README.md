# Number to Words for Go (Golang)

[![Go](https://github.com/hazzik/go-words/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/hazzik/go-words/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hazzik/go-words)](https://goreportcard.com/report/github.com/hazzik/go-words)


## Installation

```shell
go get github.com/hazzik/go-words
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/hazzik/go-words"
)

func main() {
	// prints "one hundred eleven"
	fmt.Printf("%q", words.FromInt64(111))
}
```