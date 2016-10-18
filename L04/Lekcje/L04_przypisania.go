package main

import "fmt"

func main() {
	a := 'a'
	b := 'ą'
	var c int32 = 100
	show(a)
	show(b)
	show(c)
}

func show(a int32) {
	fmt.Printf("%[1]c %[1]U %[1]Xhex %[1]vdec %[1]T\n", a)
}
