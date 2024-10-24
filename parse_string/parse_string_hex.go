package parse_string

// ParseStringDecimal erwartet einen String, der eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringHexadecimal(s string) int {
	result := 0
	for _, c := range s {
		d := ParseDigit(string(c))
		if d == -1 {
			return -1
		}
		result = result*16 + d
	}
	return result
}

// HINT
// Verwenden Sie die Funktion ParseDigit, um die Ziffern des Strings zu bestimmen.
