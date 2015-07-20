package models

import (
	"regexp"
	"unicode"
)

var (
	usernameMatcher *regexp.Regexp
	emailMatcher    *regexp.Regexp
	nameMatcher     *regexp.Regexp
	collegeMatcher  *regexp.Regexp
)

func init() {
	usernameMatcher = regexp.MustCompile("^[\\w\\d_\\.]*$")
	emailMatcher = regexp.MustCompile("^[\\w\\d_\\.]{2,}\\@[\\w\\d]*\\.[\\w]{2,4}(\\.[\\w]{2,4})?$")
	nameMatcher = regexp.MustCompile("^([\\w]{1,}\\s?){1,}$")
	collegeMatcher = regexp.MustCompile("^([\\w-]{1,}\\s?){1,}$")
}

func CheckUserName(username string) bool {
	if len(username) < 6 {
		return false
	}
	return usernameMatcher.Match([]byte(username))
}

func CheckEmail(email string) bool {
	return emailMatcher.Match([]byte(email))
}

func CheckPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	return matchPassword(password)
}

func CheckName(name string) bool {
	return nameMatcher.Match([]byte(name))
}

func CheckCollege(college string) bool {
	return collegeMatcher.Match([]byte(college))
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
