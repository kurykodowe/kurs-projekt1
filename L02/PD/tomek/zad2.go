package main

import "fmt"

func main() {
	explain_bit_operations(2, 3)
}

func explain_bit_operations(left uint8, right uint8) {
	print_bit_operation(left, right, "&", left & right)
	print_bit_operation(left, right, "|", left | right)
	print_bit_operation(left, right, "^", left ^ right)
	print_bit_operation(left, right, "&^", left &^ right)
	print_shift_operation(left, right, "<<", left << right)
	print_shift_operation(left, right, ">>", left >> right)
}

func print_bit_operation(left uint8, right uint8, operator string, result uint8) {
	fmt.Printf(
		"%[1]v %[3]v %[2]v =\n" +
			"%08[1]b %[3]v\n" +
			"%08[2]b =\n" +
			"%08[4]b =\n" +
			"%[4]v\n\n", left, right, operator, result)
}

func print_shift_operation(left uint8, right uint8, operator string, result uint8) {
	fmt.Printf(
		"%[1]v %[3]v %[2]v =\n" +
			"%08[1]b %[3]v %[2]v =\n" +
			"%08[4]b =\n" +
			"%[4]v\n\n", left, right, operator, result)
}