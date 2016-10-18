package main

import "fmt"

func main() {
	text := "Mydło lubi zabawę"
	for _, char := range text {
		fmt.Printf("%c\n", char)
	}
}
