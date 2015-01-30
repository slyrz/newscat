package main

import (
	"fmt"
	"github.com/slyrz/newscat/html"
	"github.com/slyrz/newscat/model"
	"github.com/slyrz/newscat/util"
	"os"
)

var highlight = util.IsTerminal(os.Stdout)

func printArticle(article *util.Article) {
	pre, pos := "", ""
	for _, text := range article.Text {
		if highlight {
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
	for _, input := range util.GetInput(os.Args[1:]) {
		if document, err := html.NewDocument(input.Data); err == nil {
			if article, err := ext.Extract(document); err == nil {
				// Extraction might miss the article heading. So if the text
				// doesn't start with a heading, use the article title as
				// opening heading.
				if !article.StartsWithHeading() && article.Title != "" {
					article.Prepend(util.Heading(article.Title))
				}
				printArticle(article)
			}
		}
		input.Data.Close()
	}
}
