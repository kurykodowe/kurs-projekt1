package main

func main() {

	for i := 1; i < 101; i = i + 1 {
		if i % 3 == 0 {
			print("\n" + "Fizz");
			if i % 5 == 0 {
				print("Buzz")
			}

		} else if i % 5 == 0 {
			print("\n" + "Buzz")
		} else {
			print("\n"); print(i)
		}
	}
}
