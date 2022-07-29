package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	userInput string
	banner    string
	output    string
)

func main() {
	// terminal command e.g.: go run . "hello" standard --output=banner.txt
	// followed by terminal command: cat <filename.txt> -e
	if len(os.Args) != 4 {
		fmt.Printf("Usage: go run . [STRING] [BANNER] [OPTION]\n")
		fmt.Println()
		fmt.Printf("Example: go run . something standard --output=<fileName.txt>\n")
		os.Exit(0)
	}

	userInput = os.Args[1]
	if userInput == "" {
		fmt.Println("String must be at least 1 character long")
		fmt.Println()
		os.Exit(0)
	}

	banner = os.Args[2]
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("Banner name not recognised. Banners: shadow, standard, thinkertoy")
		fmt.Println()
		os.Exit(0)
	}

	output = os.Args[3]
	if output == "[FILE NONE]" {
		log.Fatal("Bad flag formatting/ Incorrect fileName provided\n")
		os.Exit(0)
	}
	flagParsingMessage()
	output := checkFlags(output)
	UserInputSliced := strings.Split(string(userInput), "\\n")
	writeFile(output, AsciiArt(UserInputSliced, readFile(banner)))
}
