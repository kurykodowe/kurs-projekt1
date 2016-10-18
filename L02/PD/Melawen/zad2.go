package main

import "fmt"

func main() {
	var a, b uint8
	a = 2
	b = 3
	explain_bit_operations(a, b)
}

func explain_bit_operations(left uint8, right uint8) {
	fmt.Printf("%d & %d = \n", left, right)
	fmt.Printf("%08b &\n", left)
	fmt.Printf("%08b =\n", right)
	fmt.Printf("%08b =\n", left & right)
	fmt.Printf("%d \n\n", left & right)

	fmt.Printf("%d | %d = \n", left, right)
	fmt.Printf("%08b |\n", left)
	fmt.Printf("%08b =\n", right)
	fmt.Printf("%08b =\n", left | right)
	fmt.Printf("%d \n\n", left | right)

	fmt.Printf("%d ^ %d = \n", left, right)
	fmt.Printf("%08b ^\n", left)
	fmt.Printf("%08b =\n", right)
	fmt.Printf("%08b =\n", left ^ right)
	fmt.Printf("%d \n\n", left ^ right)

	fmt.Printf("%d &^ %d = \n", left, right)
	fmt.Printf("%08b &^\n", left)
	fmt.Printf("%08b =\n", right)
	fmt.Printf("%08b =\n", left &^ right)
	fmt.Printf("%d \n\n", left &^ right)

	fmt.Printf("%d >> %d = \n", left, right)
	fmt.Printf("%08b >>\n", left)
	fmt.Printf("%08b =\n", right)
	fmt.Printf("%08b =\n", left >> right)
	fmt.Printf("%d \n\n", left >> right)

	fmt.Printf("%d << %d = \n", left, right)
	fmt.Printf("%08b <<\n", left)
	fmt.Printf("%08b =\n", right)
	fmt.Printf("%08b =\n", left << right)
	fmt.Printf("%d \n\n", left << right)
}
