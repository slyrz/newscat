package main

import (
	"flag"
	"fmt"
	"github.com/slyrz/newscat/html"
	"github.com/slyrz/newscat/model"
	"github.com/slyrz/newscat/util"
	"io"
	"net/http"
	"os"
	"strings"
)

// Input stores the user-provided HTML data and its location.
type Input struct {
	Location string    // either file path or URL or empty if data was read from stdin
	Data     io.Reader // the HTML data (hopefully)
}

var (
	// highlight indicates whether newscat should use ANSI escape codes
	// to print headings and emphasized text in bold type. The default value of this flag
	// depends on the type of stdout - it's set to false if newscat isn't printing
	// onto a terminal.
	highlight = flag.Bool("highlight", util.IsTerminal(os.Stdout), "highlight headings and emphasized text")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `newscat [OPTION]... [PATH|URL]...

Options:
  -highlight  Use ANSI escape codes to format output.`)
	}
}

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
		if *highlight {
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
	flag.Parse()

	inputChannel := make(chan Input, 4)
	// Open all input (file paths or URLs) or read from stdin and write the
	// corresponding Input structs to the inputChannel channel.
	go func() {
		if flag.NArg() > 0 {
			for _, arg := range flag.Args() {
				if strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://") {
					if resp, err := http.Get(arg); err == nil {
						inputChannel <- Input{arg, resp.Body}
					}
				} else {
					if file, err := os.Open(arg); err == nil {
						inputChannel <- Input{arg, file}
					}
				}
			}
		} else {
			inputChannel <- Input{"", os.Stdin}
		}
		close(inputChannel)
	}()

	ext := model.NewExtractor()
	for input := range inputChannel {
		if article, err := html.NewDocument(input.Data); err == nil {
			if chunks := ext.Extract(article); len(chunks) > 0 {
				printChunks(chunks)
			}
		}
	}
}
