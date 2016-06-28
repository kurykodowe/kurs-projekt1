/**********************************/
/* ex1_HW2.go                        */
/* Drugaa praca domowa            */
/* autor:  Asia.asia              */
/* 11/06/2016                     */
/**********************************/

package main


func main() {

	for i :=1; i<101; i=i+1{
		if ((i%3==0) && (i%5==0)) {
			println("FizzBuzz")
		} else if (i%3  == 0) {
			println("Fizz")
		} else if (i%5  == 0) {
			println("Buzz")
		} else {
			println(i)
		}
			}
}
