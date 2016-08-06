package main

import "fmt"

func main() {
	if (true) {
		fmt.Println("To zobaczysz zawsze")
	}
	if (false) {
		fmt.Println("Tego nie zobaczysz nigdy")
	}
	if (true && false) {
		fmt.Println("Tego nie zobaczysz nigdy")
	}
	if (true || false) {
		fmt.Println("To zobaczysz zawsze")
	}
	if (!true) {
		fmt.Println("Tego nie zobaczysz nigdy")
	}
	if (!false) {
		fmt.Println("To zobaczysz zawsze")
	}
}
