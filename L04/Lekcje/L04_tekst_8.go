package main

import "fmt"

func main() {
	text := "Mydło lubi zabawę"
	for _, char := range []rune(text)[:5] {
		fmt.Printf("%c\n", char)
	}
}
