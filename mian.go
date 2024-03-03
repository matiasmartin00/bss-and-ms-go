package main

import (
	"fmt"
	"github.com/matiasmartin00/bss-and-ms-go/channels"
	"github.com/matiasmartin00/bss-and-ms-go/defers"
	"github.com/matiasmartin00/bss-and-ms-go/functions"
	"github.com/matiasmartin00/bss-and-ms-go/goroutines"
	"github.com/matiasmartin00/bss-and-ms-go/pointers"
	"github.com/matiasmartin00/bss-and-ms-go/slices"
	"os"
)

func main() {
	for {
		fmt.Println("Available options:")
		fmt.Println("1. Pointers")
		fmt.Println("2. Slices")
		fmt.Println("3. Defers")
		fmt.Println("4. Functions")
		fmt.Println("5. Channels")
		fmt.Println("6. Go Routines")
		fmt.Println("0. Exit")

		var option int
		fmt.Print("Chose your option: ")
		_, err := fmt.Scanln(&option)

		if err != nil {
			fmt.Println("Error reading option")
			os.Exit(1)
		}

		switch option {
		case 0:
			os.Exit(0)

		case 1:
			pointers.Execute()
		case 2:
			slices.Execute()
		case 3:
			defers.Execute()
		case 4:
			functions.Execute()
		case 5:
			channels.Execute()
		case 6:
			goroutines.Execute()
		default:
			fmt.Println("Invalid option")

		}
	}
}
