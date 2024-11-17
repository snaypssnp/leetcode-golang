package main

type Roman struct {
	value int
	key   string
}

var numerals = []Roman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) (res string) {
	for _, roman := range numerals {
		division := num / roman.value

		if division == 0 {
			continue
		}

		for i := 0; i < division; i++ {
			res += roman.key
		}

		num = num % roman.value
	}

	return
}
