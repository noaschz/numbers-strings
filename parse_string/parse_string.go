package parse_string

// ParseString erwartet einen String, der entweder eine Dezimal- oder eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseString(s string) int {
	// SOLUTION
	if len(s) > 2 && s[0:2] == "0x" {
		return ParseStringHexadecimal(s[2:])
	}
	return ParseStringDecimal(s)
	// SOLUTION_END
}
