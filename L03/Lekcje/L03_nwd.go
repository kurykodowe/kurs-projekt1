package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	if (len(os.Args) < 3) {
		fmt.Printf("Usage: %v number1 number2", os.Args[0])
		os.Exit(1)
	}

	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])

	for number1 != number2 {
		for number1 < number2 {
			number2 -= number1
		}
		for number2 < number1 {
			number1 -= number2
		}
	}

	fmt.Printf("NWD dla %v oraz %v to %v", os.Args[1], os.Args[2], number1)
}
