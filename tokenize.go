package jiebago

type token struct {
	Word  string
	Start int
	End   int
}

// Return words with position.
func Tokenize(sentence string, mode string, HMM bool) []token {
	tokens := make([]token, 0)
	start := 0
	var width int
	for word := range Cut(sentence, false, HMM) {
		if mode == "default" {
			width = len([]rune(word))
			tokens = append(tokens, token{word, start, start + width})
			start += width

		} else {
			runes := []rune(word)
			width = len(runes)
			for _, step := range []int{2, 3} {
				if width > step {
					for i := 0; i < width-step+1; i++ {
						gram := string(runes[i : i+step])
						if _, ok := Trie.Freq[gram]; ok {
							tokens = append(tokens, token{gram, start + i, start + i + step})
						}
					}
				}
			}
			tokens = append(tokens, token{word, start, start + width})
		}
	}
	return tokens
}
