package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Code by Lucas: function reading file and saving each line as a striang into a slice
func readFile(banner string) []string {
	file, err := os.Open(banner + ".txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sS := make([]string, 0)

	for scanner.Scan() {
		sS = append(sS, scanner.Text())
	}
	checkError(scanner.Err())

	return sS
}

func writeFile(outputFile string, ascii string) {
	err := ioutil.WriteFile(outputFile, []byte(ascii), 0777)
	checkError(err)

	/*content, err := ioutil.ReadFile(outputFile) //--> not required as for project ascii-art-output we type 'cat <filename.txt>' in terminal to print the file

	checkError(err)

	fmt.Println(string(content))*/
}

func checkError(e error) {
	if e != nil {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Println()
		fmt.Println("EX: go run . something standard --output=<fileName.txt>")
		os.Exit(0)
	}
}

func checkFlags(output string) string {
	WriteTo := os.Args[3][9:]
	// fmt.Printf("writeTo-->: %s\n", writeTo)
	if len(WriteTo) < 1 {
		fmt.Printf("fileName must be at least 1 character long.\n")
		fmt.Println()
		fmt.Printf("EX: go run . something standard --output=<fileName.txt>\n")
		os.Exit(0)
	} else {
		if strings.HasSuffix(WriteTo, ".txt") {
			WriteTo = os.Args[3][9:]
		} else {
			WriteTo = WriteTo + ".txt"
		}
	}
	return WriteTo
}

func flagParsingMessage() {
	// fmt.Printf("flagParsing-->: %s\n", os.Args[3][0:9])
	if os.Args[3][0:9] != "--output=" {
		fmt.Printf("Usage: go run . [STRING] [BANNER] [OPTION]\n")
		fmt.Println()
		fmt.Printf("EX: go run . something standard --output=<fileName.txt>\n")
		os.Exit(0)
	}
}
