package roman

import "strings"

// NToI erwartet eine Zahl und liefert die entsprechende Anzahl an I als String.
func NToI(n int) string {
	// TODO+
	return strings.Repeat("I", n)
}

// NToX erwartet eine Zahl und liefert die entsprechende Anzahl an X als String.
func NToX(n int) string {
	// TODO
	zehner := n / 10 % 10

	if zehner > 0 {
		return strings.Repeat("X", zehner)
	}
	return ""
}

// RomanBelow10 erwartet eine Zahl 1 <= n <= 10 und liefert die römische Schreibweise für n als String.
func RomanBelow10(n int) string {
	// TODO
	return ""
}

// RomanBelow100 erwartet eine Zahl 1 <= n <= 100 und liefert die römische Schreibweise für n als String.
func RomanBelow100(n int) string {
	// TODO
	return ""
}
