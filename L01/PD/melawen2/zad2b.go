package main

import "fmt"

func main() {
	var level = 3
	var hight = 3

	for i := 1; i <= level; i++ {
		for j := i; j < hight + i; j++ {
			for k := 1; k <= j; k++ {
				fmt.Printf("@")
			}
			fmt.Printf("\n")
		}
		hight++

	}
}
