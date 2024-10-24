package parse_string

// ParseDigit erwartet einen String von 0 bis 9 oder A bis F und liefert die zugehörige Zahl.
// Dabei gilt A=10, B=11, ..., F=15.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseDigit(digit string) int {
	// SOLUTION
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	for i, d := range digits {
		if d == digit {
			return i
		}
	}
	return -1
	// SOLUTION_END
}

// HINT
// Verwenden Sie eine Liste, um die Ziffern vorab zu definieren.
// Durchsuchen Sie die Liste, um die Ziffer zu finden.
