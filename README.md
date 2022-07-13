# Stacks
Thread-safe, generic stacks for Go.

![](https://github.com/twharmon/stacks/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/stacks)](https://goreportcard.com/report/github.com/twharmon/stacks) [![codecov](https://codecov.io/gh/twharmon/stacks/branch/main/graph/badge.svg?token=K0P59TPRAL)](https://codecov.io/gh/twharmon/stacks)

## Documentation
For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/stacks).

## Install
```bash
go get github.com/twharmon/stacks
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/twharmon/stacks"
)

func main() {
	// Create a new stack of ints
	s := stacks.New[int]()

	// Push some values
	s.Push(1, 2, 3)

	// Peek at the head
	s.Peek() // 3

	// Pop a value
	s.Pop() // 3

	// Get length of the stack
	s.Len() // 2

	// Get an ordered slice of the values in the stack
	s.Slice() // [1, 2]

	// Clear the stack
	s.Clear()
}
```

## Contribute
Make a pull request.

## Benchmarks
```
goos: darwin
goarch: arm64
pkg: github.com/twharmon/stacks
BenchmarkPush/push-10         	   74554	     15708 ns/op	   25256 B/op	      13 allocs/op
BenchmarkPush/push_pop-10     	   42314	     27693 ns/op	    2088 B/op	       9 allocs/op
```
