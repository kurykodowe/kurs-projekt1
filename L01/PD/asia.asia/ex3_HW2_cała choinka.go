/**********************************/
/* ex2_HW2_pólchoinki.go          */
/* Drugaa praca domowa            */
/* autor:  Asia.asia              */
/* 19/06/2016                     */
/**********************************/

package main

import "fmt"

func main() {
	for k := 0; k < 3; k = k + 1 {
		for j := 0; j < 3 + k; j = j + 1 {

			for l := 0; l < 6 - (j + k); l = l + 1 {
				fmt.Printf(" ")
			}
			for i := 0; i < (j + k) * 2 + 1; i = i + 1 {
				fmt.Printf("L")
			}

			fmt.Printf("\n")
		}
	}
}