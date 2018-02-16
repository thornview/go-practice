package main

import (
	"strings"
	"romanNumerals"
	"os"
	"fmt"
)

func main() {
	getInput()
	
}

func getInput() {
	fmt.Print('[R]oman to int, [I]nt to Roman, [C]alculator, [Q]uit?')
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader. ReadRune()
	if err != nil {
		fmt.Println(err)
	} 
	choice := strings.ToUpper(result)
	switch choice {
	case 'R':
		runRomanToInt()
		break
	case 'I':
		runIntToRoman()
		break
	case 'Q':
		return	
	default:
		fmt.Printf("%s is not a valid command", choice)
		getInput()
	}
}

func runRomanToInt() {
	fmt.Printf('Enter RomanNumeral: ')

}

func runIntToRoman() {
	fmt.Printf('Enter Integer: ')

}

func runCalc() {

}
