package main

import "fmt"

func main() {
	for number := 1; number < 101; number++ {
		if number % 3 == 0 {
			fmt.Print("Fizz")
		}
		if number % 5 == 0 {
			fmt.Print("Buzz")
		}
		if number % 3 != 0 && number % 5 != 0 {
			fmt.Print(number)
		}
		fmt.Println()
	}
}

