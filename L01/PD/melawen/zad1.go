package main

import "fmt"

func main() {

	for i := 1; i <= 100; i = i + 1 {
		if (i % 15 == 0) {
			fmt.Printf("FizzBuzz")
		} else if (i % 3 == 0) {
			fmt.Printf("Fizz")
		} else if (i % 5 == 0) {
			fmt.Printf("Buzz")
		} else {
			fmt.Printf("%v", i)
		}

		fmt.Printf("\n")
	}
}
