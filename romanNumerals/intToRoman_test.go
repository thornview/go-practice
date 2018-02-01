package romanNumerals

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{1971, "MCMLXXI"},
		{2018, "MMXVIII"},
	}

	for _, test := range tests {
		got, _ := intToRoman(test.input)
		if got != test.want {
			t.Errorf("intToRoman() = %v, wanted %v", got, test.want)
		}
	}
}
