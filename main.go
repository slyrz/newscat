package main

import (
	"flag"
	"fmt"
	"github.com/slyrz/newscat/html"
	"github.com/slyrz/newscat/model"
	"github.com/slyrz/newscat/util"
	"os"
)

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
	flag.Parse()
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
	ext := model.NewExtractor()
	for _, input := range util.GetInput() {
		if article, err := html.NewDocument(input.Data); err == nil {
			if chunks := ext.Extract(article); len(chunks) > 0 {
				printChunks(chunks)
			}
		}
		input.Data.Close()
	}
}
