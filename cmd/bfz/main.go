package main

import (
	"flag"
	"fmt"
)

func main() {
	targetString := flag.String("s", "", "Target string")
	shift := flag.Bool("shift", false, "Produce wordlist where each byte incremented individually")

	flag.Parse()

	if *shift {
		shiftBytes(*targetString)
	}
}

func shiftBytes(encString string) {
	for i, _ := range encString {
		r := []rune(encString)
		r[i]++
		fmt.Println(string(r))
	}
}
