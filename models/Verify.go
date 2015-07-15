package models

import (
	"regexp"
	"unicode"
)

var (
	usernameMatcher *regexp.Regexp
)

func init() {
	usernameMatcher, _ = regexp.Compile("^[\\w\\d_]*$")
}

func CheckUserName(username string) bool {
	if len(username) < 6 {
		return false
	}
	return usernameMatcher.Match([]byte(username))
}

func CheckPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	return matchPassword(password)
}

func matchPassword(password string) bool {
	var (
		letter bool
		number bool
	)
	for _, r := range password {
		if unicode.IsLetter(r) {
			letter = true
		} else if unicode.IsNumber(r) {
			number = true
		} else {
			return false
		}
	}
	return letter && number
}
