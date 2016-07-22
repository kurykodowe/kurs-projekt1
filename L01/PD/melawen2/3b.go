package main

import "fmt"

func main() {
	const TREE_MODULES_AMOUNT = 3
	const MODULE_HEIGHT_BASE = 3
	const LINE_WIDTH_BASE = 1

	const LINE_MARGIN_BASE = 2 * TREE_MODULES_AMOUNT

	for module_index := 0; module_index < TREE_MODULES_AMOUNT; module_index++ {
		var module_height = MODULE_HEIGHT_BASE + module_index

		for line_index := 0; line_index < module_height; line_index++ {
			var line_width = LINE_WIDTH_BASE + module_index + line_index
			line_margin := LINE_MARGIN_BASE - module_index - line_index

			for k := 0; k < line_margin; k++ {
				fmt.Printf(" ")
			}

			for l := 0; l < 2 * line_width; l++ {
				fmt.Printf("@")
			}

			fmt.Printf("\n")
		}
	}
}
