package natural

// NumberStringGreater20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringGreater20(n int) string {
	if n%100 < 20 && n%100 > 0 {
		return DigitString100(n/100) + NumberStringBelow20(n%100)
	}
	return DigitString100(n/100) + DigitString1(n%10) + DigitString10((n%100)/10)
}

// HINT
// Verwenden Sie Modulo- und Divisions-Operationen, um die Ziffern der Zahl zu bestimmen.
// Verwenden Sie die Funktionen DigitString1, DigitString10 und DigitString100,
// um die Ziffern in Strings umzuwandeln.
