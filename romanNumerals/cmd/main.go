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
	mux.Handle("/", http.FileServer(http.Dir("../static")))
	mux.HandleFunc("/romToInt", romToInt)
	mux.HandleFunc("/intToRom", intToRom)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func romToInt(w http.ResponseWriter, req *http.Request) {
	romanNum := req.URL.Query().Get("roman")
	int, _ := romanNumerals.RomanToInt(romanNum)
	fmt.Fprintf(w, "%v", int)
}

func intToRom(w http.ResponseWriter, req *http.Request) {
	input := req.URL.Query().Get("number")
	num, _ := strconv.Atoi(input)
	romanNum, _ := romanNumerals.IntToRoman(num)
	fmt.Fprintf(w, "%s", romanNum)
}
