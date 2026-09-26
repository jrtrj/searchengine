package main

import (
	"fmt"

	"github.com/jrtrj/searchengine/internal/index"
	"github.com/jrtrj/searchengine/internal/model"
)

func main() {
	doc1 := model.Document{
		Id:    1,
		Title: "search",
		Body:  "go go search",
	}
	doc2 := model.Document{
		Id:    2,
		Title: "server",
		Body:  "go web server",
	}
	doc3 := model.Document{
		Id:    3,
		Title: "search",
		Body:  "go search engine",
	}
	index := index.NewIndex()
	index.Add(&doc1)
	index.Add(&doc2)
	index.Add(&doc3)

	if docs, err := index.Search("engine"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(docs)
	}
}
