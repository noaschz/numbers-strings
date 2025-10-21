package natural

// NumberStringBelow20 erwartet eine Zahl <= 20 und liefert den zugehörigen String.
func NumberStringBelow20(n int) string {
	// TODO
	var result string

	switch n {
	case 0:
		return "null"
	case 1:
		return "eins"
	case 2:
		return "zwei"
	case 3:
		return "drei"
	case 4:
		return "vier"
	case 5:
		return "fünf"
	case 6:
		return "sechs"
	case 7:
		return "sieben"
	case 8:
		return "acht"
	case 9:
		return "neun"
	case 10:
		return "zehn"
	case 11:
		return "elf"
	case 12:
		return "zwölf"
	case 20:
		return "zwanzig"
	default:
		einer := n - 10
		switch einer {
		case 3:
			result = "drei"
		case 4:
			result = "vier"
		case 5:
			result = "fünf"
		case 6:
			result = "sech"
		case 7:
			result = "sieb"
		case 8:
			result = "acht"
		case 9:
			result = "neun"
		}
		return result + "zehn"
	}
}
