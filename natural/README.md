# Zahl in gesprochene Sprache umwandeln

Die Aufgabe hier ist, eine Funktion zu schreiben,
die eine Zahl als `int` erwartet und die einen String
zurückgibt, der die Zahl ausschreibt.

## Beispiele

**Beispiele (Normalfall):**

* 143 -> "einhundertdreiundvierzig"
* 53 -> "dreiundfünfzig"
* 7 -> "sieben"
* 258 -> "zweihundertachtundfünfzig"

**Beispiele (Sonderfälle):**

* 12 -> "zwölf"
* 20 -> "zwanzig"
* 131 -> "einhunderteinunddreißig"
* 1 -> "eins"
* 0 -> "null"
* 258756 -> "zweihundertachtundfünfzigtausendsiebenhundertsechsundfünfzig"

## Hinweis zur Vorgehensweise

Die Aufgabe ist in eine Reihe von Teilaufgaben zerlegt, die immer komplexere
Aspekte des Problem behandeln und aufeinander aufbauen.

* Beginnen Sie mit den Aufgaben `digit_string_...`.
  Diese bestimmen zunächst die Schreibweise für Einer-, Zehner- und Hunterterstellen.
* Die Aufgaben `number_string_below_20` und `number_string_greater_20`
  nutzen dann die `digit_string_...`- Funktionen, um Zahlen zusammenzusetzen.
  Die Zahlen bis 20 werden separat behandelt, weil sie sehr viele Sonderfälle enthalten.
* Die Aufgaben `number_string_n_digits` setzen dann alles zu einer Gesamtlösung zusammen.
