package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Brak drugiego argumentu")
		os.Exit(1)
	}
	number, _ := strconv.Atoi(os.Args[1])
	if (number % 2 == 0) {
		fmt.Printf("%v: Liczba parzysta\n", number)
	} else {
		fmt.Printf("%v: Liczba nieparzysta\n", number)
	}
}
