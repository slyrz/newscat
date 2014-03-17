package main

import (
	"fmt"
	"github.com/slyrz/newscat/html"
	"github.com/slyrz/newscat/model"
	"os"
)

func main() {
	svm, err := model.Open(model.Path)
	if err != nil {
		panic(err)
	}
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
		for i, feature := range model.Features(doc) {
			if svm.Predict(&feature) {
				fmt.Println(doc.Chunks[i].Text)
			}
		}
	}
}
