package validator

import "regexp"

// IsLetter var
var HasLowwerCaseLetter = regexp.MustCompile(`[a-z]`).MatchString
var HasUpperCaseLetter = regexp.MustCompile(`[A-Z]`).MatchString
var HasDigit = regexp.MustCompile(`[0-9]`).MatchString
