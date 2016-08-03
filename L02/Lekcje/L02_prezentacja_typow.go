package main

import "fmt"

func main() {
	fmt.Printf("%T\n", 29)
	fmt.Printf("%08b bin\n", 29)
	fmt.Printf("%08d dec\n", 29)
	fmt.Printf("%08o oct\n", 29)
	fmt.Printf("%08x hex\n", 29)
	fmt.Printf("%b %T\n", 1.0, 1.0)
	fmt.Printf("%e %T\n", 1.0, 1.0)
	fmt.Printf("%b %T\n", 1i, 1i)
}
