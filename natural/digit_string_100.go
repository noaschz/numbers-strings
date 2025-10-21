package natural

// DigitString100 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Hunderter-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreihundert" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString100(digit int) string {
	// TODO
	var result string
	var hundert string

	switch digit {
	case 0:
		result = ""
	case 1:
		result = "ein"
	case 2:
		result = "zwei"
	case 3:
		result = "drei"
	case 4:
		result = "vier"
	case 5:
		result = "fünf"
	case 6:
		result = "sechs"
	case 7:
		result = "sieben"
	case 8:
		result = "acht"
	case 9:
		result = "neun"
	}

	if digit != 0 {
		hundert = "hundert"
	}

	return result + hundert
}
