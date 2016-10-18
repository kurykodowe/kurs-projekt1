package main

import "fmt"

func main() {

	//fmt.Println("9223372036854775807+1 =", 9223372036854775807+1)
	/*wynik:
	# command-line-arguments
	.\zad1.go:6: constant 9223372036854775808 overflows int
	Process finished with exit code 2
	*/

	//fmt.Println("-9223372036854775808-1 =", -9223372036854775808-1)

	/* wynik:
	# command-line-arguments
	.\zad1.go:10: constant -9223372036854775809 overflows int
	Process finished with exit code 2
	*/

	//jednak w przypadku deklaracji zmiennej:
	var x int = 9223372036854775807
	var y int
	y = x + 1
	fmt.Printf("%d \n", x)
	fmt.Printf("%d \n", y)

	//i odejmowanie:
	var z int
	z = y - 1
	fmt.Printf("%d \n", z)

	/*UZASADNIENIE:
	Przy deklaracji zmiennej widać, ze "przekręciliśmy stos", czyli z powodu braku miejsca rozpoczęło się naliczanie od nowa.
	Jak w grze w państwa i miasta jak dojeżdżamy do z to zaczynamy liczyć od nowa, tak komputer przypisał kolejną wolną liczbę.
	W przypadku, jeśli gdzieś chcielibyśmy użyć takiej zmiennej wywołałoby to trudny do wykrycia błąd (aplikacja nie zwraca błędu, jedynie otrzyane wyniki znacznie różnią się od oczekiwanych).

	 */

}