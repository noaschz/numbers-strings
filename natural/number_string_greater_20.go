package natural

// NumberStringGreater20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringGreater20(n int) string {

	var result string

	hunderter := n / 100 // 0..9
	rest := n % 100      // 0..99

	if hunderter > 0 {
		// z. B. 1 -> "einhundert", 2 -> "zweihundert", ...
		result += DigitString100(hunderter)
	}

	// Nichts hinter den Hundertern?
	if rest == 0 {
		return result // z. B. 100, 200, 300 ...
	}

	// Sonderfälle 0..20 (inkl. 10–19): direkt aus Below20 anhängen
	if rest <= 20 {
		return result + NumberStringBelow20(rest)
	}

	// Ab 21: "EinerUNDZehner"
	zehner := rest / 10 // 2..9
	einer := rest % 10  // 0..9

	// Einer zuerst (liefert bei dir z. B. "einund", "zweiund", ..., "" bei 0)
	if einer > 0 {
		result += DigitString1(einer)
	}

	// Dann die Zehner (z. B. 2 -> "zwanzig", 3 -> "dreißig", 4 -> "vierzig", ...)
	result += DigitString10(zehner)

	return result
}
