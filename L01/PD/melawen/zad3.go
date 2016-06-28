package main

import "fmt"

func main() {
	var rozmiar_wciecia = 6

	for i := 1; i <= 3; i++ {
		for j := i; j < 2 * i + 2; j++ {
			for k := j + 1; k < 2 * i + rozmiar_wciecia; k++ {
				fmt.Printf(" ")
			}
			for l := 1; l < 2 * j + 1; l++ {
				fmt.Printf("@")
			}

			fmt.Printf("\n")
		}
		rozmiar_wciecia = rozmiar_wciecia - 2
	}
}
