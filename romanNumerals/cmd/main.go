package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/romToInt", http.HandlerFunc(romToInt))
	mux.Handle("/intToRom", http.HandlerFunc(intToRom))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func romToInt(w http.ResponseWriter, req *http.Request) {
	romanNum := req.URL.Query().Get("roman")
	fmt.Fprintf(w, "{requested: %v}", romanNum)
}

func intToRom(w http.ResponseWriter, req *http.Request) {
	num := req.URL.Query().Get("int")
	fmt.Fprintf(w, "you asked for: %v", num)
}
