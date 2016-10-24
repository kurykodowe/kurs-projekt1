package main

import "fmt"

func main() {
	text := "Mydło lubi zabawę"
	for _, char := range string([]rune(text)[:5]) {
		fmt.Printf("%c\n", char)
	}
}
