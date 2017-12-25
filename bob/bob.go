/*
Package bob simulates the bob, a teenager who has typical
replies in every conversation.

1. Bob answers 'Sure' if he is asked a question.
2. He answers 'Whoa, chill out!' if he is shouted on(this is detected with the help of presence of all Capital letters in the conversation)
3. He says, 'Fine. Be that way!' if he is actually not told anything.
4. If he is asked a question, while shouting, he replies 'Calm down, I know what I\'m doing!'
5. If he is told anything else, he answers 'Whatever'
*/
package bob

import (
	"strings"
)

// Hey decides what kind of remark is made to Bob
func Hey(remark string) string {
	remark = sanitizeString(remark)

	// if the remark is an empty statement
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	q := isQuestion(remark)
	s := isScream(remark)
	cl := containsLetter(remark)
	if cl && q && s {
		return "Calm down, I know what I'm doing!"
	}
	if q {
		return "Sure."
	}
	if cl && s {
		return "Whoa, chill out!"
	}
	if cl {
		return "Whatever."
	}
	return "Whatever."
}

// sanitizeString removes the leading and trailing blank characters.
func sanitizeString(remark string) string {
	remark = strings.TrimSpace(remark)
	return remark
}

// isQuestion checks if the remark is a question
func isQuestion(remark string) bool {
	if string(remark[len(remark)-1]) == "?" {
		return true
	}
	return false
}

// isScream checks if bob is being shouted on
func isScream(remark string) bool {
	upperRemark := strings.ToUpper(remark)
	if remark == upperRemark {
		return true
	}
	return false
}

// containsLetter checks if the string contains anything other than letters and blank spaces.
func containsLetter(remark string) bool {
	for _, r := range remark {
		if (r > 'A' && r < 'Z') || (r > 'a' && r < 'z') {
			return true
		}
	}
	return false
}
