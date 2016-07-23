/**********************************/
/* ex2_HW2_pólchoinki.go          */
/* Drugaa praca domowa            */
/* autor:  Asia.asia              */
/* 11/06/2016                     */
/**********************************/

package main

import "fmt"

func main() {
	for k := 0; k < 3; k = k + 1 {
		for j := 0; j < 3 + k; j = j + 1 {

			for i := 0; i < j + k + 1; i = i + 1 {
				fmt.Printf("*")
			}
			fmt.Printf("\n")
		}
	}
}