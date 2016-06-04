package main

import "fmt"

func main() {
	for i:=1;i<=3;i++{
		for j:=i;j<2*i+2;j++{
			for k:=1;k<j+1;k++{
				fmt.Printf("@")
			}
			fmt.Printf("\n")
		}
	}
}
