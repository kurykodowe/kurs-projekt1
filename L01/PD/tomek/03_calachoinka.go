package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3 + i; j++ {
			for k := 6; k > i + j; k-- {
				fmt.Print(" ")
			}
			for k := 0; k < 2*(1 + i + j)-1; k++ {
				fmt.Print("@")
			}
			fmt.Println()
		}
	}
}
