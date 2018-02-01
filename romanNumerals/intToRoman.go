package romanNumerals

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

var letterMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var romanNumber string

func intToRoman() {
	number, err := strconv.Atoi(os.Args[1])
	if err != nil || number >= 4000 || number < 1 {
		fmt.Printf("%v is an invalid value.  Input must be an integer between 1 - 4000", number)
	} else {
		convertIntToRoman(number)
	}
}

func convertIntToRoman(number int) {
	// Convert map into a slice so it can be sorted
	var values []int
	for value := range letterMap {
		values = append(values, value)
	}

	// Sort it from largest to smallest
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	// Now we can iterate over the slice
	for _, value := range values {
		mapIntToLetters(value)
	}
}

func mapIntToLetters(value int) {
	fmt.Printf("I see the number %d", value)
}
