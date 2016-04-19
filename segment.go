package searcher

// 超简单的分词  列出分词所有可能性
func SimpleSegment(str string) (ret []string) {
	var w = []rune(str)
	var key = map[string]bool{}
	for i := 0; i != len(w); i++ {
		for j := 0; j+i < len(w); j++ {
			key[string(w[j:j+i+1])] = true
		}
	}
	for k, _ := range key {
		ret = append(ret, k)
	}
	return
}
