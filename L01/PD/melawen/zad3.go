package main

import "fmt"

func main() {
	var a = 6

	for i := 1; i <= 3; i++ {
		for j := i; j < 2 * i + 2; j++ {
			for k := j + 1; k < 2 * i + a; k++ {
				fmt.Printf(" ")
			}
			for l := 1; l < j + 1; l++ {
				fmt.Printf("@")
			}
			for l := 1; l < j + 1; l++ {
				fmt.Printf("@")
			}

			fmt.Printf("\n")
		}
		a = a - 2
	}
}
