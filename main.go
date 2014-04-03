package main

import (
	"fmt"
	"github.com/slyrz/newscat/html"
	"github.com/slyrz/newscat/model"
	"os"
)

func printChunks(chunks []*html.Chunk) {
	var last *html.Chunk = nil
	var delim string = ""
	var pre string = ""
	var pos string = ""
	for _, chunk := range chunks {
		// If the last chunk and the current chunk share the same HTML block element,
		// we seperate them by a space character. If they are in different blocks, we
		// use two newline characters.
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
	ext := model.NewExtractor()
	for _, arg := range os.Args[1:] {
		file, err := os.Open(arg)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		doc, err := html.NewDocument(file)
		if err != nil {
			panic(err)
		}
		printChunks(ext.Extract(doc))
	}
}
