package main

import (
	"os"
	"math"
	"fmt"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		println("Provide an integer to perform verification")
		return
	}

	number, _ := strconv.Atoi(os.Args[1])

	if number < 2 {
		println("Value too small for prime verification: " + strconv.Itoa(number))
		return
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number % i == 0 {
			fmt.Printf("%v is composite\n", number)
			return
		}
	}
	fmt.Printf("%v is prime\n", number)
}
