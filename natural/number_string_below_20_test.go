package natural

import "fmt"

func ExampleNumberStringBelow20() {
	for i := 0; i < 20; i++ {
		fmt.Println(NumberStringBelow20(i))
	}

	// Output:
	// null
	// eins
	// zwei
	// drei
	// vier
	// fünf
	// sechs
	// sieben
	// acht
	// neun
	// zehn
	// elf
	// zwölf
	// dreizehn
	// vierzehn
	// fünfzehn
	// sechzehn
	// siebzehn
	// achtzehn
	// neunzehn
}
