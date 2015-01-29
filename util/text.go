package util

import (
	"bytes"
	"strings"
	"unicode"
)

type Text struct {
	Words     int
	Sentences int
	// Unexported fields.
	buffer bytes.Buffer
	words  *Stringset
}

func NewText() *Text {
	text := new(Text)
	text.words = NewStringset()
	return text
}

// isWord returns true if the passed text seems to be an actual word and not
// clutter like email addresses or URLs.
func isWord(text string) bool {
	letters := 0
	for _, rune := range text {
		if unicode.IsLetter(rune) {
			letters += 1
		}
	}
	// A "real" word must have more than 2 characters (sorry "it")
	// and is only allowed to contain at most 2 non-letter characters.
	return (len(text) > 2) && (letters >= (len(text) - 2))
}

func (t *Text) WriteText(s *Text) {
        t.WriteString(s.String())
}

func (t *Text) WriteString(s string) {
	// If buffer contains text, write a space first to avoid joining words
	// accidentally.
	needSpace := t.buffer.Len() > 0

	// Split sentence into words. Count number of words and sentences and add
	// each word to the string set, so we can compare texts based on the number
	// of identical words they contain.
	for _, word := range strings.Fields(s) {
		if needSpace {
			t.buffer.WriteRune(' ')
		}
		t.buffer.WriteString(word)
		// Check if the current word is a "real" word.
		if isWord(word) {
			t.words.Add(word)
			t.Words += 1
		}
		// Check if the current text part ends a sentence.
		switch word[len(word)-1] {
		case '!', '.', '?':
			t.Sentences += 1
		}
		needSpace = true
	}
}

// Calculate a word-based similarity to a given text. This function returns
// values between [0,1], where zero means the texts share no words and one
// means the text have all words in common. This function is fuzzy.
func (t *Text) Similarity(u *Text) float32 {
	a := t.words.Common(&u.words.Bitset)
	b := t.words.Len()
	c := u.words.Len()
	if b > c {
		return float32(a) / float32(b)
	} else {
		return float32(a) / float32(c)
	}
}

func (t *Text) String() string {
	return t.buffer.String()
}

func (t *Text) Len() int {
	return t.buffer.Len()
}
