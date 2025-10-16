package fizzbuzz

import "fmt"

//Fizz gibt die Zahlen von 1 bis 30 aus, ersetzt dabei
//jede durch 3 teilbare Zahl durch "fizz",
//jede durch 5 teilbare Zahl durch "buzz" und
//jede durch 3 und 5 teilbare Zahl durch "fizzbuzz"

func FizzCompact() {
	for i := 1; i <= 30; i++ {
		//wenn i weder durch 3 noch 5 teilbar
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}

		if i%3 == 0 {
			fmt.Print("fizz")
		}
		if i%5 == 0 {
			fmt.Print("buzz")
		}

		fmt.Println() // sonst gib i aus
	}
}
