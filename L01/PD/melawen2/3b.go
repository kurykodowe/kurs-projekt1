package main

import "fmt"

func main() {
	var level = 3
	var hight = 3
	var a = 2*level

	for i:=1;i<=level;i++{
		for j:=i;j<hight+i;j++{
			for k:=j+1;k<2*i+a;k++{
				fmt.Printf(" ")
			}
			for l:=1;l<j+1;l++{
				fmt.Printf("@")
			}
			for l:=1;l<j+1;l++{
				fmt.Printf("@")
			}

			fmt.Printf("\n")
		}
		hight++
		a=a-2
	}
}
