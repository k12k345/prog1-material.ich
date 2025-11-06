package matrices

import "fmt"

func Example_matrix() {
	m1 := [][]int{{1, 2}, {3, 4}}

	fmt.Println(m1[0][0])
	//Output:
	//1
}

//Printmatrixngib eine Zeile für Zeile aus
func Example_PrintMatrix() {
	m1 := [][]int{{1, 2}, {3, 4}}
	PrintMatrix(m1)
	//Output:
	// 1 2
	// 3 4
}
