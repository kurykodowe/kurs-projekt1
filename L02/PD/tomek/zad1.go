package main

import "fmt"

func main() {
	a := 9223372036854775807
	fmt.Printf("%t %x %b %v\n", a, a, a, a + 1)
	a = -9223372036854775808
	fmt.Printf("%t %x %b %v\n", a, a, a, a - 1)
}
