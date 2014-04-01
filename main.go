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
	for _, chunk := range chunks {
		switch {
		case last == nil:
			delim = ""
		case last.Block != chunk.Block:
			delim = "\n\n"
		case last.Block == chunk.Block:
			delim = " "
		}
		fmt.Printf("%s%s", delim, chunk.Text)
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
