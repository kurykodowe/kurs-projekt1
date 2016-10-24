package main

/*odp:
program rzucał cały czas błędem: "Brak drugiego argumentu"
 */
import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	var i int
	if (len(os.Args) < 2) {
		fmt.Println("Brak drugiego argumentu")
		os.Exit(1)
	}


	for i=1;i<len(os.Args);i++{
		number, err := strconv.Atoi(os.Args[i])

		if(err!=nil){
			fmt.Printf("Sorry, ale "+os.Args[i]+" to nie liczba, jak chcesz sprawdzić jej parzystość?")
		}else if(number % 2 == 0) {
			fmt.Printf("%v: Liczba parzysta", number)
		} else {
			fmt.Printf("%v: Liczba nieparzysta", number)
		}
		fmt.Println()
	}

}
