package main

import (
	"github.com/matiasmartin00/bss-and-ms-go/defers"
	"github.com/matiasmartin00/bss-and-ms-go/pointers"
	"github.com/matiasmartin00/bss-and-ms-go/slices"
)

func main() {
	// pointers 1.5.3
	pointers.Execute()

	// slices 1.5.6
	slices.Execute()

	// defers 1.6.5
	defers.Execute()
}
