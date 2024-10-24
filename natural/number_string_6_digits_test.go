package natural

import "fmt"

func ExampleNumberString6Digits() {
	fmt.Println(NumberString6Digits(113642))
	fmt.Println(NumberString6Digits(100000))
	fmt.Println(NumberString6Digits(100013))
	fmt.Println(NumberString6Digits(786))
	fmt.Println(NumberString6Digits(16384))
	fmt.Println(NumberString6Digits(0))
	fmt.Println(NumberString6Digits(103))

	// Output:
	// einhundertdreizehntausendsechshundertzweiundvierzig
	// einhunderttausend
	// einhunderttausenddreizehn
	// siebenhundertsechsundachtzig
	// sechzehntausenddreihundertvierundachtzig
	// null
	// einhundertdrei
}
