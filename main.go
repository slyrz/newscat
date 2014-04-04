package main

import (
	"fmt"
	"github.com/slyrz/newscat/html"
	"github.com/slyrz/newscat/model"
	"io"
	"net/http"
	"os"
	"strings"
)

func printChunks(chunks []*html.Chunk) {
	var last *html.Chunk = nil
	for _, chunk := range chunks {
		delim, pre, pos := "", "", ""
		// If the last chunk and the current chunk are part of the same HTML block,
		// we seperate them by a space character. If they are in different blocks,
		// we use two newline characters.
		switch {
		case last == nil:
			delim = ""
		case last.Block != chunk.Block:
			delim = "\n\n"
		case last.Block == chunk.Block:
			delim = " "
		}
		// Use bold font for headings, emphasized and bold text.
		switch chunk.Base.Data {
		case "h1", "h2", "h3", "h4", "h5", "h6", "em", "strong", "b":
			pre, pos = "\x1b[39;1m", "\x1b[0m"
		default:
			pre, pos = "", ""
		}
		fmt.Printf("%s%s%s%s", delim, pre, chunk.Text, pos)
		last = chunk
	}
	fmt.Println()
}

func main() {
	inputData := make(chan io.Reader, 4)
	// Open all input (either paths / URLs passed as command line arguments or
	// read from stdin) and write it to inputData channel.
	go func() {
		if args := os.Args[1:]; len(args) > 0 {
			for _, arg := range args {
				if strings.HasPrefix(arg, "http://") {
					if resp, err := http.Get(arg); err == nil {
						inputData <- resp.Body
					}
				} else {
					if file, err := os.Open(arg); err == nil {
						inputData <- file
					}
				}
			}
		} else {
			inputData <- os.Stdin
		}
		close(inputData)
	}()

	ext := model.NewExtractor()
	// Read input from inputData channel and perform content extraction.
	for data := range inputData {
		// TODO: Warn if parsing Document failed.
		if doc, err := html.NewDocument(data); err == nil {
			// TODO: Print warning if no chunks were extracted.
			if chunks := ext.Extract(doc); len(chunks) > 0 {
				printChunks(chunks)
			}
		}
	}
}
