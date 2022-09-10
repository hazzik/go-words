# Number to Words for Go (Golang)

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