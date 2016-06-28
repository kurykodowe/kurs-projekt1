/**********************************/
/* ex2_HW2_pólchoinki.go          */
/* Drugaa praca domowa            */
/* autor:  Asia.asia              */
/* 11/06/2016                     */
/**********************************/

package main

import "fmt"

func main() {
	for k:= 0; k < 5; k = k + 1 {
		for j := 1; j < 4; j = j + 1 {

			for i :=1; i < j+k; i = i + 1 {
				fmt.Printf("*")
			}
			fmt.Printf("\n")
		}
	}
}