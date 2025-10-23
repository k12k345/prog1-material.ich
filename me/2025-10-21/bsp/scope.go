package varnames

import "fmt"

func Example_scopes_loops() {
	x := 42
	y := 23

	for x := 0; x < 5; x++ {
		fmt.Println(x)
		fmt.Println(y)

	}
	fmt.Println(x)
}

//scopes variable verliert gültigkeit bei schleife/funktionen, immer im scope deklarieren
// variablen kann man neu deklarieren bei scope
//in vorlesung 2025-10-21 varnames
//zu ihm funtions_test.go
//&&
//scopes_test.go
//ordner varnames

//UND CONTAINS
