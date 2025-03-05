package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	targetString := flag.String("s", "", "Target string")
	inc := flag.Bool("inc", false, "Produce wordlist where each byte incremented individually")
	dec := flag.Bool("dec", false, "Produce wordlist where each byte decremented individually")

	flag.Parse()

	if *targetString == "" {
		fmt.Println("-s required")
		os.Exit(2)
	}

	if *inc {
		incBytes(*targetString)
	}
	if *dec {
		decBytes(*targetString)
	}
}

func incBytes(encString string) {
	for i, _ := range encString {
		r := []rune(encString)
		r[i]++
		fmt.Println(string(r))
	}
}

func decBytes(encString string) {
	for i, _ := range encString {
		r := []rune(encString)
		r[i]--
		fmt.Println(string(r))
	}
}
