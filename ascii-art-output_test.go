/*
To run the test file, un-comment rows 32 to 36 in 'formulas.go' file, save file and type 'go test . ' in the terminal.
*/
package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

func ReadTestFile() []string { // type 'go test .' in terminal to run this test file
	data, err := ioutil.ReadFile("test-output.txt")
	if err != nil {
		panic(err)
	}

	contentString := string(data)

	contentSplit := strings.Split(contentString, "#")

	return contentSplit
}

func TestAsciiArt(t *testing.T) {
	testCases := []string{
		"First\\nTest",
		"hello",
		"123 -> #$%",
		"432 -> #$%&@",
		"There",
		"123 -> \"#$%@",
		"2 you",
		"Testing long output!",
	}

	expectedOutput := ReadTestFile()

	for index, testCase := range testCases {
		var cmd *exec.Cmd

		switch testCase {
		case "First\\nTest":
			cmd = exec.Command("go", "run", ".", testCase, "shadow", "--output=test00.txt")
		case "hello":
			cmd = exec.Command("go", "run", ".", testCase, "standard", "--output=test01.txt")
		case "123 -> #$%":
			cmd = exec.Command("go", "run", ".", testCase, "standard", "--output=test02.txt")
		case "432 -> #$%&@":
			cmd = exec.Command("go", "run", ".", testCase, "shadow", "--output=test03.txt")
		case "There":
			cmd = exec.Command("go", "run", ".", testCase, "shadow", "--output=test04.txt")
		case "123 -> \"#$%@":
			cmd = exec.Command("go", "run", ".", testCase, "thinkertoy", "--output=test05.txt")
		case "2 you":
			cmd = exec.Command("go", "run", ".", testCase, "thinkertoy", "--output=test06.txt")
		case "Testing long output!":
			cmd = exec.Command("go", "run", ".", testCase, "standard", "--output=test07.txt")
		}
		output, _ := cmd.Output()
		if string(output) != expectedOutput[index] {
			t.Errorf("\nTest fails when given case:\n\t\"%s\","+"\nThe test should show:\n%s\nInstead it shows:\n%s", testCase, expectedOutput[index], output)
		}
	}
}
