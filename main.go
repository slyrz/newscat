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

func printArticle(article *util.Article) {
	pre, pos := "", ""
	for _, text := range article.Text {
		if *highlight {
			switch text.(type) {
			case util.Heading:
				pre, pos = "\x1b[1m", "\x1b[0m"
			case util.Paragraph:
				pre, pos = "", ""
			}
		}
		fmt.Printf("%s%s%s\n\n", pre, text, pos)
	}
}

func main() {
	ext := model.NewExtractor()
	for _, input := range util.GetInput() {
		if article, err := html.NewDocument(input.Data); err == nil {
			if article, err := ext.Extract(article); err == nil {
				printArticle(article)
			}
		}
		input.Data.Close()
	}
}
