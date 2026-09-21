package index

import (
	"github.com/jrtrj/searchengine/internal/model"
	"github.com/jrtrj/searchengine/internal/tokenizer"
)

type InvertedIndex struct {
	index map[string][]model.Posting
}

func NewIndex() *InvertedIndex {
	return &InvertedIndex{
		index: make(map[string][]model.Posting),
	}
}

// Inverted Index
func (idx *InvertedIndex) Add(doc *model.Document) {
	tokenCounts := tokenizer.Tokenize(doc.Body)
	for token, count := range tokenCounts {
		posting := model.Posting{
			DocId: doc.Id,
			Count: count,
		}
		idx.index[token] = append(idx.index[token], posting)
	}
}
