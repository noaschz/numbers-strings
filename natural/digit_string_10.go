package natural

// DigitString10 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Zehner-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreißig" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString10(digit int) string {
	// TODO
	var result string

	switch digit {
	case 0:
		result = ""
	case 1:
		result = ""
	case 2:
		result = "zwanzig"
	case 3:
		result = "dreißig"
	case 4:
		result = "vierzig"
	case 5:
		result = "fünfzig"
	case 6:
		result = "sechzig"
	case 7:
		result = "siebzig"
	case 8:
		result = "achtzig"
	case 9:
		result = "neunzig"
	}

	return result
}
