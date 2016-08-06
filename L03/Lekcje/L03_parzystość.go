package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	number, _ := strconv.Atoi(os.Args[1])
	if (number % 2 == 0) {
		fmt.Println("Liczba parzysta")
	} else {
		fmt.Println("Liczba nieparzysta")
	}
}
