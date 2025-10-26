package natural

// NumberStringBelow20 erwartet 0 <= n <= 20 und liefert den zugehörigen String.
func NumberStringBelow20(n int) string {
	if n < 0 || n > 20 {
		return "Keine gültige Zahl"
	}

	words := []string{
		"null",     // 0
		"eins",     // 1
		"zwei",     // 2
		"drei",     // 3
		"vier",     // 4
		"fünf",     // 5
		"sechs",    // 6
		"sieben",   // 7
		"acht",     // 8
		"neun",     // 9
		"zehn",     // 10
		"elf",      // 11
		"zwölf",    // 12
		"dreizehn", // 13
		"vierzehn", // 14
		"fünfzehn", // 15
		"sechzehn", // 16
		"siebzehn", // 17
		"achtzehn", // 18
		"neunzehn", // 19
		"zwanzig",  // 20
	}
	return words[n]
}
