package natural

// NumberString6Digits erwartet eine Zahl 0 <= n <= 999999 und liefert den zugehörigen String.
func NumberString6Digits(n int) string {
	//TODO
	if n < 0 || n > 999999 {
		return "Keine gültige Zahl"
	}

	var result string
	if n == 0 {
		return "null"
	}

	low := n % 1000
	high := n / 1000

	if high > 0 {
		result += NumberString3Digits(high) + "tausend"
	}
	if low > 0 {
		result += NumberString3Digits(low)
	}
	return result
}
