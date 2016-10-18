package main

import "fmt"

func main() {
	text := "Mydło lubi zabawę"
	printed_counter := 0
	for _, char := range text {
		fmt.Printf("%c\n", char)
		printed_counter += 1
		if printed_counter >= 5 {
			break
		}
	}
}
