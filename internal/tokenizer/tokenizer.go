package tokenizer

import (
	"strings"
)

func Tokenize(input string) map[string]int {
	term := make(map[string]int)
	tokens := strings.Fields(input)
	for _, val := range tokens {
		term[strings.ToLower(val)]++
	}
	return term
}