package listsearch

import "fmt"

//sucht x in l und liedert die position
//falls x nixht in l vorkommt, eird -1 geliefert

func Find(l []int, x int) int {
	//TODO
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
	//Output:
	//2
	//-1
}

//sucht x in l und liefert eine liste
// mit allen vorkommen von c und l
func FindAll(l []int, x int) []int {
	result := []int{}

	for i := 0; i < len(l); i++ {
		result = append(result, i)
		if l[i] == x {
			return result
		}

		return result
	}
}
func ExampleFind() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23}

	pos1 := Find(l1, 42)
	pos2 := Find(l1, 200)

	fmt.Println(pos1)
	fmt.Println(pos2)

	// Output:
	//[2]
	//[]
	//[]
}
