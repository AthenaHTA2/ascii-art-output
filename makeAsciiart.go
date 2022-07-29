/*
building a result ascii-art
iterating over each string in a slice and adding new line after each element of text
for each character it concatenates corresponding element from 'art'slice to final string
*/
package main

import (
	"log"
)

func AsciiArt(text []string, art []string) (final string) {
	for _, str := range text {
		for i := 0; i < 8; i++ {
			for _, v := range str {
				final += art[GetCharStartingIndex(int(v), i)]
			}
			final += "\n"
		}
		final += "\n"
	}
	return
}

/*function returning specific index of an ascii art character
using the arithmetic progression logic
*/

func GetCharStartingIndex(char int, counter int) int {
	if char > 126 {
		log.Fatal("Characters out of ASCII range not implemented.")
	}
	return (1 + (char-32)*8) + (char - 32) + counter
}
