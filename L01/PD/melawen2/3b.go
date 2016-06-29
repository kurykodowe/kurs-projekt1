package main

import "fmt"

func main() {
	var level = 3
	var height = 3
	var a = 2 * level

	for i := 1; i <= level; i++ {
		for j := i; j < height + i; j++ {
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
		height++
		a = a - 2
	}
}
