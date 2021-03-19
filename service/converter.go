package service

import "fmt"

type CorrespondingNumerals struct {
	arabic int
	roman  string
}

var Register = []CorrespondingNumerals{
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
	{1, "I"}}

func TransformArabicToRoman(number int) (result string, err error) {
	if number < 1 || number > 3999 {
		err = fmt.Errorf("the number %d is beyond the roman numerals boundaries", number)
		return
	}

	for _, pair := range Register {
		for ; number >= pair.arabic; number -= pair.arabic {
			result += pair.roman
		}
	}

	return
}
