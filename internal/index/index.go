package index

import (
	"fmt"
	"sort"

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

func (idx *InvertedIndex) getTokenDocId(token string) []int {
	postingList := idx.index[token]
	var result []int
	for _, posting := range postingList {
		result = append(result, posting.DocId)
	}
	return result
}

func (idx *InvertedIndex) Search(query string) ([]int,error) {
	queryToken := tokenizer.Tokenize(query)
	if len(queryToken) == 0 {
        return nil, fmt.Errorf("query cannot be empty")
  }
	resultSet := make(map[int]struct{})
	for token, _ := range queryToken {
		docIdSet := idx.getTokenDocId(token)
		for _,docId := range docIdSet {
			resultSet[docId] = struct{}{}
		}
	}

	result := make([]int, 0, len(resultSet))
	for docId, _ := range resultSet {
		result = append(result, docId)
	}
	sort.Ints(result)
	return result,nil
}
