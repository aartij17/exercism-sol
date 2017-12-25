/*
package acronym takes in a single/multiword string and returns an abbreviation.
This string may have phrases containing delimiters.
*/
package acronym

import (
	"strings"
)

// Abbreviate returns the abbreviation for a given string
func Abbreviate(s string) string {
	// 1. Get the words in the phrase
	// 2. Go over each word. Extract the letters till they are uppercase.
	// 3. Concantenate all these letters.
	tokens := strings.FieldsFunc(s, Split)
	finalString := getAllCapLetters(tokens)
	return finalString
}

// Split returns a list of tokens obtained in the presence of multiple delimiters in a string.
func Split(r rune) bool {
	return r == ' ' || r == '-' || r == ':'
}

// getAllCapLetters gets all the capital letters from the tokens obtained
func getAllCapLetters(tokens []string) string {
	finalString := ""
	for _, token := range tokens {
		finalString = finalString + strings.ToUpper(string(token[0]))
	}
	return finalString
}
