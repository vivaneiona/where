package main

import (
	"fmt"

	"github.com/vivaneiona/where"
)

func main() {
	fmt.Printf("where.Has(\"us-east-1\"): %v\n", where.Has("us-east-1"))
	fmt.Printf("where.Has(\"invalid-region\"): %v\n", where.Has("invalid-region"))
}
