# rot13

A simple command line utility which performs rot13 encryption. Project intended for learning the go toolkit/standards.

## WIP

This utility can perform rot13 encoding on ascii text while passing through whitespace and punctuation. It has the following optional flags:

* -caps forces the ciphertext output to be all caps (as is traditional!)
* -read can take a path to a text file to be encrypted.
* -write can take a path to a text file to save the ciphertext into.
* -h or -help lists all flags.

Example valid program calls:

```
> rot13 hello world
> rot13 -caps hello world
> rot13 -read="testfile.txt"
> rot13 -write="output.txt" hello world
> rot13 -read="testfile.txt" -write="output.txt"
```

## Description

*rot13* is a standard shift cipher which shifts letters 13 places through the alphabet - ensuring that the same function can be used to encode and decode text. *rot13* is most often used as a technique for disguising spoilers in communication methods that don't have a built-in method for hiding text. I made this utility because it was a reasonable way to learn some of the go toolkit and also because having a command line utility for this seemed fun!
