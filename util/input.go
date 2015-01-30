package util

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// Input stores the user-provided data and its origin.
type Input struct {
	Origin string        // either file path or URL or empty if data was read from stdin
	Data   io.ReadCloser // the HTML data (hopefully)
}

func GetInput(args []string) []Input {
	result := make([]Input, 0)
	if len(args) > 0 {
		for _, arg := range args {
			if strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://") {
				if resp, err := http.Get(arg); err == nil {
					result = append(result, Input{arg, resp.Body})
				}
			} else {
				if file, err := os.Open(arg); err == nil {
					result = append(result, Input{arg, file})
				}
			}
		}
	} else {
		result = append(result, Input{"", os.Stdin})
	}
	return result
}
