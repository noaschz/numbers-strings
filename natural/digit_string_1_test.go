package natural

import "fmt"

func ExampleDigitString1() {
	fmt.Println(DigitString1(0))
	fmt.Println(DigitString1(1))
	fmt.Println(DigitString1(2))
	fmt.Println(DigitString1(3))
	fmt.Println(DigitString1(4))
	fmt.Println(DigitString1(5))
	fmt.Println(DigitString1(6))
	fmt.Println(DigitString1(7))
	fmt.Println(DigitString1(8))
	fmt.Println(DigitString1(9))

	// Output:
	//
	// einund
	// zweiund
	// dreiund
	// vierund
	// fünfund
	// sechsund
	// siebenund
	// achtund
	// neunund
}
