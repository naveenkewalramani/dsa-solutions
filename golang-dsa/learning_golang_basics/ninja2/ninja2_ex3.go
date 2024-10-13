package main

import (
	"fmt"
)

type customInt int

// UntypedInt is a constant
const (
	UntypedInt               = 29
	CustomTypedInt customInt = 31
)

func main() {
	fmt.Println(UntypedInt)
	fmt.Println(CustomTypedInt)
}
