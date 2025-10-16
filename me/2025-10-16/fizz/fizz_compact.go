package fizzbuzz

import "fmt"

//Fizz gibt die Zahlen von 1 bis 30 aus, ersetzt dabei
//jede durch 3 teilbare Zahl durch "fizz",
//jede durch 5 teilbare Zahl durch "buzz" und
//jede durch 3 und 5 teilbare Zahl durch "fizzbuzz"

func FizzCompact() {
	for i := 1; i <= 30; i++ {

		// wenn i durch 3 und 5 teilbar ist, gib "fizzbuzz" aus
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")

		} else if
		// wenn i durch 3 teilbar ist, gib "fizz" aus
		i%3 == 0 { //% ist modula, kein rest bei rechnung
			fmt.Println("fizz")
		} else if
		// wenn i durch 5 teilbar ist, gib "buzz" aus
		i%5 == 0 {
			fmt.Println("buzz")
			// fmt.Printf("%10d, %v\n", i, "buzz")
		} else {
			fmt.Println(i)
		}
		// sonst gib i aus
	}

}
