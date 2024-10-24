package roman

import "fmt"

func ExampleNToI() {
	fmt.Println(NToI(1))
	fmt.Println(NToI(2))
	fmt.Println(NToI(3))

	// Output:
	// I
	// II
	// III
}

func ExampleRomanBelow10() {
	fmt.Println(RomanBelow10(0))
	fmt.Println(RomanBelow10(1))
	fmt.Println(RomanBelow10(2))
	fmt.Println(RomanBelow10(3))
	fmt.Println(RomanBelow10(4))
	fmt.Println(RomanBelow10(5))
	fmt.Println(RomanBelow10(6))
	fmt.Println(RomanBelow10(7))
	fmt.Println(RomanBelow10(8))
	fmt.Println(RomanBelow10(9))
	fmt.Println(RomanBelow10(10))

	// Output:
	//
	// I
	// II
	// III
	// IV
	// V
	// VI
	// VII
	// VIII
	// IX
	// X
}

func ExampleRomanBelow100_from11to20() {
	fmt.Println(RomanBelow100(11))
	fmt.Println(RomanBelow100(12))
	fmt.Println(RomanBelow100(13))
	fmt.Println(RomanBelow100(14))
	fmt.Println(RomanBelow100(15))
	fmt.Println(RomanBelow100(16))
	fmt.Println(RomanBelow100(17))
	fmt.Println(RomanBelow100(18))
	fmt.Println(RomanBelow100(19))
	fmt.Println(RomanBelow100(20))

	// Output:
	// XI
	// XII
	// XIII
	// XIV
	// XV
	// XVI
	// XVII
	// XVIII
	// XIX
	// XX
}

func ExampleRomanBelow100_from21to30() {
	fmt.Println(RomanBelow100(21))
	fmt.Println(RomanBelow100(22))
	fmt.Println(RomanBelow100(23))
	fmt.Println(RomanBelow100(24))
	fmt.Println(RomanBelow100(25))
	fmt.Println(RomanBelow100(26))
	fmt.Println(RomanBelow100(27))
	fmt.Println(RomanBelow100(28))
	fmt.Println(RomanBelow100(29))
	fmt.Println(RomanBelow100(30))

	// Output:
	// XXI
	// XXII
	// XXIII
	// XXIV
	// XXV
	// XXVI
	// XXVII
	// XXVIII
	// XXIX
	// XXX
}

func ExampleRomanBelow100_from31to40() {
	fmt.Println(RomanBelow100(31))
	fmt.Println(RomanBelow100(32))
	fmt.Println(RomanBelow100(33))
	fmt.Println(RomanBelow100(34))
	fmt.Println(RomanBelow100(35))
	fmt.Println(RomanBelow100(36))
	fmt.Println(RomanBelow100(37))
	fmt.Println(RomanBelow100(38))
	fmt.Println(RomanBelow100(39))
	fmt.Println(RomanBelow100(40))

	// Output:
	// XXXI
	// XXXII
	// XXXIII
	// XXXIV
	// XXXV
	// XXXVI
	// XXXVII
	// XXXVIII
	// XXXIX
	// XL
}

func ExampleRomanBelow100_from41to50() {
	fmt.Println(RomanBelow100(41))
	fmt.Println(RomanBelow100(42))
	fmt.Println(RomanBelow100(43))
	fmt.Println(RomanBelow100(44))
	fmt.Println(RomanBelow100(45))
	fmt.Println(RomanBelow100(46))
	fmt.Println(RomanBelow100(47))
	fmt.Println(RomanBelow100(48))
	fmt.Println(RomanBelow100(49))
	fmt.Println(RomanBelow100(50))

	// Output:
	// XLI
	// XLII
	// XLIII
	// XLIV
	// XLV
	// XLVI
	// XLVII
	// XLVIII
	// XLIX
	// L
}

func ExampleRomanBelow100_from51to60() {
	fmt.Println(RomanBelow100(51))
	fmt.Println(RomanBelow100(52))
	fmt.Println(RomanBelow100(53))
	fmt.Println(RomanBelow100(54))
	fmt.Println(RomanBelow100(55))
	fmt.Println(RomanBelow100(56))
	fmt.Println(RomanBelow100(57))
	fmt.Println(RomanBelow100(58))
	fmt.Println(RomanBelow100(59))
	fmt.Println(RomanBelow100(60))

	// Output:
	// LI
	// LII
	// LIII
	// LIV
	// LV
	// LVI
	// LVII
	// LVIII
	// LIX
	// LX
}

func ExampleRomanBelow100_from61to70() {
	fmt.Println(RomanBelow100(61))
	fmt.Println(RomanBelow100(62))
	fmt.Println(RomanBelow100(63))
	fmt.Println(RomanBelow100(64))
	fmt.Println(RomanBelow100(65))
	fmt.Println(RomanBelow100(66))
	fmt.Println(RomanBelow100(67))
	fmt.Println(RomanBelow100(68))
	fmt.Println(RomanBelow100(69))
	fmt.Println(RomanBelow100(70))

	// Output:
	// LXI
	// LXII
	// LXIII
	// LXIV
	// LXV
	// LXVI
	// LXVII
	// LXVIII
	// LXIX
	// LXX
}

func ExampleRomanBelow100_from71to80() {
	fmt.Println(RomanBelow100(71))
	fmt.Println(RomanBelow100(72))
	fmt.Println(RomanBelow100(73))
	fmt.Println(RomanBelow100(74))
	fmt.Println(RomanBelow100(75))
	fmt.Println(RomanBelow100(76))
	fmt.Println(RomanBelow100(77))
	fmt.Println(RomanBelow100(78))
	fmt.Println(RomanBelow100(79))
	fmt.Println(RomanBelow100(80))

	// Output:
	// LXXI
	// LXXII
	// LXXIII
	// LXXIV
	// LXXV
	// LXXVI
	// LXXVII
	// LXXVIII
	// LXXIX
	// LXXX
}

func ExampleRomanBelow100_from81to90() {
	fmt.Println(RomanBelow100(81))
	fmt.Println(RomanBelow100(82))
	fmt.Println(RomanBelow100(83))
	fmt.Println(RomanBelow100(84))
	fmt.Println(RomanBelow100(85))
	fmt.Println(RomanBelow100(86))
	fmt.Println(RomanBelow100(87))
	fmt.Println(RomanBelow100(88))
	fmt.Println(RomanBelow100(89))
	fmt.Println(RomanBelow100(90))

	// Output:
	// LXXXI
	// LXXXII
	// LXXXIII
	// LXXXIV
	// LXXXV
	// LXXXVI
	// LXXXVII
	// LXXXVIII
	// LXXXIX
	// XC
}

func ExampleRomanBelow100_from91to100() {
	fmt.Println(RomanBelow100(91))
	fmt.Println(RomanBelow100(92))
	fmt.Println(RomanBelow100(93))
	fmt.Println(RomanBelow100(94))
	fmt.Println(RomanBelow100(95))
	fmt.Println(RomanBelow100(96))
	fmt.Println(RomanBelow100(97))
	fmt.Println(RomanBelow100(98))
	fmt.Println(RomanBelow100(99))
	fmt.Println(RomanBelow100(100))

	// Output:
	// XCI
	// XCII
	// XCIII
	// XCIV
	// XCV
	// XCVI
	// XCVII
	// XCVIII
	// XCIX
	// C
}
