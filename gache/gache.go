package main

import (
	"os"

	"github.com/k-oguma/go-cache-benchmarks/gache/lib"
)

func main() {
	lib.Gache(os.Args[1], os.Args[2])
}

