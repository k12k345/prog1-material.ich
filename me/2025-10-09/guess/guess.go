package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// my_number := 42
	my_number := rand.Intn(100) + 1

	for i := 0; i < 3; i = i + 1 {
		guess := ReadNumber()
		if NumberGood(guess, my_number) {
			fmt.Println("Das war richtig:-)")
			return // Programm beenden, dann zu else, geht aber ohne else wenn return
			// } else { //Klammern so
		}
		fmt.Printf("%d war nicht korrekt!\n", guess) //alt Gr
	}

	fmt.Println("Game Over")
}

// ReadNumber liefert uns ein int-Zahl
func ReadNumber() int {
	var n int // alternativ n:=0

	fmt.Print("Bitte gib eine Zahl ein:")
	fmt.Scan(&n)

	return n
}

// NumberGood prüft ob x eine zufällig gewählte Zahl zwischen 1 und 100 ist
func NumberGood(x int, n int) bool {

	return x == n

}
