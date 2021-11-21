package leetcode

type WordDictionary struct {
	word  []string
	total int
}

func Constructor() WordDictionary {
	var w WordDictionary
	w.word = make([]string, 0)
	w.total = 0
	return w
}

func (this *WordDictionary) AddWord(word string) {
	this.word = append(this.word, word)
	this.total++
	return
}

func (this *WordDictionary) Search(word string) bool {
	if this.total == 0 {
		return false
	}

	i := 0
	for {
	T:
		if this.word[i] == word {
			return true
		}

		if len(this.word[i]) != len(word) {
			i++
			if i >= this.total {
				return false
			}
			goto T
		}

		for k := range this.word[i] {
			if word[k] != (this.word[i])[k] && word[k] != '.' {
				i++
				if i >= this.total {
					return false
				}
				goto T
			}
		}

		return true

	}
	//return false
}
