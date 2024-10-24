# Arabische in römische Zahlen umwandeln

Bei dieser Aufgabe geht es darum, eine Funktion zu schreiben,
die eine Zahl als `int` erwartet und die einen String
zurückgibt, der die Zahl in römischen Ziffern ausschreibt.

## Beispiele

* 1 -> "I"
* 2 -> "II"
* 4 -> "IV"
* 9 -> "IX"
* 38 -> "XXXVIII"
* 42 -> "XLII"
* 99 -> "XCIX"

## Hinweise

Die vorgebenen Tests prüfen alle Zahlen von 0 bis 100.
Für größere Zahlen gibt es keine Tests, aber die Funktion
könnte nach dem gleichen Schema wie in den hier erarbeiteten
Lösungen erweitert werden.

Die Tests sind in 10er-Blöcke unterteilt.
Beginnen Sie mit der Funktion `RomanBelow10` und erfüllen Sie die entsprechenden Tests.
Sobald `RomanBelow10` funktioniert, können Sie sich an `RomanBelow100` wagen.
Dabei sollten Sie die Tests für die Zahlenblöcke 11-20, 21-31, ...
nacheinander zum Laufen bringen und Ihre Funktion entsprechend erweitern.
