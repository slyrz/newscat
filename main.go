package main

import (
	"fmt"
	"github.com/slyrz/newscat/html"
	"github.com/slyrz/newscat/model"
	"github.com/slyrz/newscat/util"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	highlight = util.IsTerminal(os.Stdout)
)

func printChunks(chunks []*html.Chunk) {
	var last *html.Chunk = nil
	for _, chunk := range chunks {
		delim, pre, pos := "", "", ""
		// If the last chunk and the current chunk are part of the same HTML block,
		// separate them by a space character. Otherwise use two newline characters
		// to create a new paragraph.
		if last != nil {
			switch {
			case last.Block != chunk.Block:
				delim = "\n\n"
			case last.Block == chunk.Block:
				delim = " "
			}
		}
		if highlight {
			// Print headings and emphasized text bold.
			switch chunk.Base.Data {
			case "h1", "h2", "h3", "h4", "h5", "h6", "em", "strong", "b":
				pre, pos = "\x1b[1m", "\x1b[0m"
			default:
				pre, pos = "", ""
			}
		}
		fmt.Printf("%s%s%s%s", delim, pre, chunk.Text, pos)
		last = chunk
	}
	fmt.Println()
}

func main() {
	inputData := make(chan io.Reader, 4)
	// Open all input (file paths or URLs) or read from stdin and write the
	// corresponding io.Readers to the inputData channel.
	go func() {
		if args := os.Args[1:]; len(args) > 0 {
			for _, arg := range args {
				if strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://") {
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
		// TODO: Warn if parsing document failed.
		if doc, err := html.NewArticle(data); err == nil {
			// TODO: Print warning if no chunks were extracted.
			if chunks := ext.Extract(doc); len(chunks) > 0 {
				printChunks(chunks)
			}
		}
	}
}
