package main

import (
	"os"
	"fmt"
	"strconv"
)


func main(){
	//fmt.Println("Podaj liczbę\n")

	if (len(os.Args) != 2) {
		fmt.Println("Brak drugiego argumentu")
		os.Exit(1)
	}

	number1,_ := strconv.Atoi(os.Args[1])


	if (number1<2){
		fmt.Println("Liczba musi być większa niż 2")
		os.Exit(1)
	}
	if (number1%1!=0){
		fmt.Println("Podaj liczbę całkowitą")
		os.Exit(2)
	}

	if(pierwsza(number1)){
		fmt.Println("To jest liczba pierwsza")
	}else{
		fmt.Println("To nie jest liczba pierwsza")
	}
}

func pierwsza (number int) bool{

	var i int

	for i=2; i*i<=number;i++  {
		if (number%i==0){
			return false
		}
		
	}
	return true
}