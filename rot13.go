package main

import (
	"flag"
	"fmt"
	"os"
	s "strings"
)

func main() {

	// Set up command line flags
	// -caps is a boolean flag to toggle ciphertext being presented in all caps
	// default: false
	var capToggle bool
	flag.BoolVar(&capToggle, "caps", false, "Should ciphertext be presented in all caps? Default: false")

	// -read gets a path to a text file which the utility loads to encrypt.
	// default: ""
	var readPath string
	flag.StringVar(&readPath, "read", "", "Path to a text file to read. Default: empty")

	// -write gets a path to a text file which the utility will output the encrypted message to.
	// default: ""
	var writePath string
	flag.StringVar(&writePath, "write", "", "Path to a text file to write the result to. Default: empty")

	flag.Parse()

	// check if a read file path was provided. if not, pull from command line arguments.

	var plaintext string = ""

	if readPath == "" {

		// get plaintext by converting arguments from array to single string

		for _, arg := range flag.Args() {
			plaintext += arg
			plaintext += " "
		}

	} else {

		// get plaintext by reading entire file in

		data, err := os.ReadFile(readPath)
		check(err)
		plaintext += string(data)

	}

	// if no plaintext was provided display generic help prompts,
	// otherwise run encryption

	if plaintext == "" {

		fmt.Println("rot13 is a command line utility that performs rot13 encryption on text.")
		fmt.Println("Use it to encrypt and decrypt text.")
		fmt.Println("Example usage: 'rot13 hello world'")
		fmt.Println("Type 'rot13 -h' for optional flags.")

	} else {

		// go through plaintext one rune at a time and parse each character

		var ciphertext string = ""
		for _, rn := range plaintext {
			ciphertext += runeShifter(rn)
		}

		// capitalise the ciphertext if requested

		if capToggle {
			ciphertext = s.ToUpper(ciphertext)
		}

		// check if a write file path was provided. if not, output to command line.

		if writePath == "" {

			// print to command line

			fmt.Println(ciphertext)

		} else {

			// convert ciphertext to byte data and write to the file

			outputData := []byte(ciphertext)
			err := os.WriteFile(writePath, outputData, 0644)
			check(err)

		}

	}

}

func runeShifter(rn rune) string {
	// we need to treat capital and lower case letters differently
	// so first we need to determine which is happening here

	if rn >= 65 && rn <= 90 {

		// capitals - 65 - 90

		rn += 13
		if rn > 90 {
			rn -= 26
		}

	} else if rn >= 97 && rn <= 122 {

		// lower case - 97 - 122

		rn += 13
		if rn > 122 {
			rn -= 26
		}

	}

	// other (punc, whitespace, etc)
	// these pass through without change

	return string(rn)

}

func check(e error) {
	// simple error parser

	if e != nil {
		panic(e)
	}

}
