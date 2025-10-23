package searching

func BinFind2(l []int, x int) int {
	links := 0
	rechts := len(l)
	for rechts > links {
		mitte := (rechts+links)/2 + links
		if x == l[mitte] {

			return mitte
		}
		if x < l[mitte] {

			rechts = mitte
		} else {

			links += mitte + 1
		}
	}
	return -1
}
