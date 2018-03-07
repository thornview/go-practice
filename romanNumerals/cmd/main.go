package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/thornview/romanNumerals"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderWebPage)
	mux.HandleFunc("/romToInt", romToInt)
	mux.HandleFunc("/intToRom", intToRom)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func romToInt(w http.ResponseWriter, req *http.Request) {
	romanNum := req.URL.Query().Get("roman")
	int, _ := romanNumerals.RomanToInt(romanNum)
	fmt.Fprintf(w, "{%s = %v}", romanNum, int)
}

func intToRom(w http.ResponseWriter, req *http.Request) {
	input := req.URL.Query().Get("int")
	num, _ := strconv.Atoi(input)
	romanNum, _ := romanNumerals.IntToRoman(num)
	fmt.Fprintf(w, "%v = %s", num, romanNum)
}

func renderWebPage(w http.ResponseWriter, req *http.Request) {
	page := "<html><head><title>Go Roman Calculator</title></head><body><h1>Roman Numeral Calculator</h1></body></html>"
	fmt.Fprintf(w, "you asked for: %s", page)
}
