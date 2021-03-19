package service

import (
	"os"
	"strings"
	"testing"

	"github.com/MaksimYudenko/andersen/utils"
)

var randomValues, testValues []CorrespondingNumerals

func beforeTest() {
	testValues = createTestValues()
}

func TestMain(m *testing.M) {
	beforeTest()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRomanNumerals(t *testing.T) {
	t.Run("normal roman numbers comparison", func(t *testing.T) {
		for i, value := range randomValues {
			expected := randomValues[i].roman
			actual, err := TransformArabicToRoman(value.arabic)

			utils.Ok(t, err)
			utils.LogValues(t, expected, actual)
			utils.Equals(t, expected, actual)
		}
	})

	t.Run("invalid roman numbers comparison", func(t *testing.T) {
		for i, value := range randomValues {
			expected := randomValues[i].roman
			actual, err := TransformArabicToRoman(value.arabic + 1)
			isEqual := strings.EqualFold(expected, actual)

			utils.Ok(t, err)
			utils.LogValues(t, false, isEqual)
			utils.Equals(t, false, isEqual)
		}
	})
}

func BenchmarkRomanNumerals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testValues {
			_, err := TransformArabicToRoman(test.arabic)
			utils.Ok(b, err)
		}
	}
}

func createTestValues() []CorrespondingNumerals {
	randomValues = []CorrespondingNumerals{
		{69, "LXIX"},
		{444, "CDXLIV"},
		{986, "CMLXXXVI"},
		{1429, "MCDXXIX"},
		{2021, "MMXXI"},
	}

	return append(Register, randomValues...)
}
