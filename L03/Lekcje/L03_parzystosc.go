package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	number, _ := strconv.Atoi(os.Args[1])
	if (number % 2 == 0) {
		fmt.Printf("%v: Liczba parzysta\n", number)
	} else {
		fmt.Printf("%v: Liczba nieparzysta\n", number)
	}
}
