package util

import (
	"regexp"
	"strings"
)

type Regex struct {
	*regexp.Regexp
}

// NewRegex creates a new Regex struct matching the given pattern.
func NewRegex(pattern string) *Regex {
	// Create a case insensitive regular expression which matches all given
	// arguments.
	return &Regex{regexp.MustCompile(pattern)}
}

// NewWordMatchRegex creates a new Regex struct matching any of the words
// passed as arguments.
func NewRegexFromWords(words ...string) *Regex {
	// Create a case insensitive regular expression which matches all given
	// arguments.
	return NewRegex("(?i)" + strings.Join(words, "|"))
}

// In returns true if a match can be found in text.
func (r *Regex) In(text string) bool {
	return r.FindStringIndex(text) != nil
}
