package romanNumerals

import (
	"reflect"
	"testing"
)

func TestConvertRomanToInt(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{"XXIV", []int{10, 10, 1, 5}},
		{"MDCLXVI", []int{1000, 500, 100, 50, 10, 5, 1}},
	}
	for _, test := range tests {
		got := convertRomanToInt(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("convertRomanToInt() = %v, want %v", got, test.want)
		}
	}
}

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{10, 1, 10}, 19},
		{[]int{1000, 100, 1000, 50, 10, 10, 1, 5}, 1974},
	}

	for _, test := range tests {
		got := calculateTotal(test.input)
		if got != test.want {
			t.Errorf("calculateTotal() = %v, want %v", got, test.want)
		}
	}
}

func TestIsValidRomanNumeral(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
		err   error
	}{
		{[]int{0, 1, 1}, false, nil},
		{[]int{10, 1, 5}, true, nil},
	}
	for _, test := range tests {
		got, _ := isValidRomanNumeral(test.input)
		if got != test.want {
			t.Errorf("isValidRomanNumeral() = %v, want %v", got, test.want)
		}
	}
}
