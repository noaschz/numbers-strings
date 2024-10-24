package parse_string

// ParseStringDecimal erwartet einen String, der eine Dezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringDecimal(s string) int {
	// SOLUTION
	result := 0
	for _, c := range s {
		d := ParseDigit(string(c))
		if d == -1 {
			return -1
		}
		result = result*10 + d
	}
	return result
	// SOLUTION_END
}

// HINT
// Verwenden Sie die Funktion ParseDigit, um die Ziffern des Strings zu bestimmen.
