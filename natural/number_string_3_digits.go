package natural

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {
	// TODO
	if n < 0 || n > 999 {
		return "Keine gültige Zahl"
	}

	if n <= 20 {
		return NumberStringBelow20(n)
	}

	var result string

	if n >= 21 {
		einer := n % 10
		zehner := (n / 10) % 10
		hunderter := (n / 100) % 10

		if hunderter > 0 {
			result += DigitString100(hunderter)
		}

		if einer > 0 {
			result += DigitString1(einer)
		}

		if zehner > 0 {
			result += DigitString10(zehner)
		}
	}
	return result
}
