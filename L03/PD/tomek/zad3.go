package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Brak liczby do sprawdzenia parzystości")
		os.Exit(1)
	}

	for _, num_str := range os.Args[1:] {
		number, err := strconv.Atoi(num_str);
		if err != nil {
			fmt.Printf("%v: Błąd konwersji: %v\n", num_str, err.Error())
			continue
		}
		if (number % 2 == 0) {
			fmt.Printf("%v: Liczba parzysta\n", number)
		} else {
			fmt.Printf("%v: Liczba nieparzysta\n", number)
		}
	}
}
