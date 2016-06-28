package main

import "fmt"

func main() {
	var level = 3
	var height = 3

	for i := 1; i <= level; i++ {
		for j := i; j < height + i; j++ {
			for k := 1; k <= j; k++ {
				fmt.Printf("@")
			}
			fmt.Printf("\n")
		}
		height++

	}
}
