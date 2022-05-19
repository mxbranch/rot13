# rot13

A simple command line utility which performs rot13 encryption. Project intended for learning the go toolkit/standards.

## WIP

This utility currently:

* Can perform rot13 encoding on ascii text while passing through whitespace and punctuation.
* Has an optional flag -caps to output the ciphertext in caps (as is traditional!)

In the future the following features will be added:

* Reading/writing files

## Description

*rot13* is a standard shift cipher which shifts letters 13 places through the alphabet - ensuring that the same function can be used to encode and decode text. *rot13* is most often used as a technique for disguising spoilers in communication methods that don't have a built-in method for hiding text. I made this utility because it was a reasonable way to learn some of the go toolkit and also because having a command line utility for this seemed fun!
