// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if empty(remark) {
		return "Fine. Be that way!"
	} else if allCapital(remark) && containLetter(remark) && endsWithQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if allCapital(remark) && containLetter(remark) {
		return "Whoa, chill out!"
	} else if endsWithQuestion(remark) {
		return "Sure."
	} else {
		return "Whatever."
	}
}

func allCapital(remark string) bool {
	return strings.ToUpper(remark) == remark
}

func containLetter(remark string) bool {
	for _, r := range remark {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func endsWithQuestion(remark string) bool {
	return strings.HasSuffix(strings.TrimSpace(remark), "?")
}

func empty(remark string) bool {
	return len(strings.TrimSpace(remark)) == 0
}
