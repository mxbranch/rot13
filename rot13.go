package main

import (
	"flag"
	"fmt"
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
	flag.StringVar(&readPath, "read", "", "NOT IMPLEMENTED: Path to a text file to read. Default: empty")

	// -write gets a path to a text file which the utility will output the encrypted message to.
	// default: ""
	var writePath string
	flag.StringVar(&writePath, "write", "", "NOT IMPLEMENTED: Path to a text file to write the result to. Default: empty")

	flag.Parse()

	// get plaintext by converting arguments from array to single string

	var plaintext string = ""
	for _, arg := range flag.Args() {
		plaintext += arg
		plaintext += " "
	}

	// go through plaintext one rune at a time and parse each character

	var ciphertext string = ""
	for _, rn := range plaintext {
		ciphertext += runeShifter(rn)
	}

	// capitalise the ciphertext if requested

	if capToggle {
		ciphertext = s.ToUpper(ciphertext)
	}

	fmt.Println(ciphertext)

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
