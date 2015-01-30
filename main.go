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
		if article, err := html.NewDocument(input.Data); err == nil {
			if article, err := ext.Extract(article); err == nil {
				printArticle(article)
			}
		}
		input.Data.Close()
	}
}
