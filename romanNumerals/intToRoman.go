package romanNumerals

import (
	"fmt"
	"os"
	"strconv"
)

var letterMap = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman() {
	number, err := strconv.Atoi(os.Args[1])
	if err == nil && (number >= 4000 || number < 1) {
		fmt.Printf("Thank you")
	} else {
		fmt.Printf("%v is an invalid value.  Input must be an integer between 1 - 4000", number)
	}
}
