package searching

// Vergleiche x mit dem Element in der Mitte.
//wenn x == l[mitte] dann fertig
//wenn x< l[mitte] dann suche im linken teil weiter
//wenn x> l[mitte] dann suche im rechten teil weiter
func BinFind1(l []int, x int) int {
	links := 0
	for {
		mitte := len(l) / 2
		if x == l[mitte] {

			return mitte + links
		}
		if x < l[mitte] {
			// lasse nur den linken teil übrig
			// alles von 0 bis exklusiv mitte
			l = l[0:mitte]
		} else {
			// wenn x>l[mitte] dann suche im rechten teil
			// lasse nur den rechten teil übrig
			// alles von mitte+1 bis ende
			l = l[mitte+1 : len(l)]
			links += mitte + 1
		}
		return -1
	}
}
