package main

import "fmt"	

func main() {
	guess := Readnumber	()
	fmt.Println(guess)
	guess = Readnumber	() // kein : weil gint es schon
	fmt.Println(guess)
	guess = Readnumber	()
	fmt.Println(guess)

}

// Readnumber liefert uns ein int-Zahl
func Readnumber() int {
	var n int // alternatin n:=0

	fmt.Print("Bitte gib eine Zahl ein:")
	fmt.Scan(&n)	

	return n
}
