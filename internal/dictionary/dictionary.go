package dictionary

type dictionaryEntry struct {
	prefix int  // index of the previous dictionary code
	ch     byte // character appended to the prefix
}

func InitDictionary() ([]dictionaryEntry, map[[2]int]int) {
	dict := make([]dictionaryEntry, 256)
	lookup := make(map[[2]int]int, 256)
	// Initialize single-byte entries
	for i := 0; i < 256; i++ {
		dict[i] = dictionaryEntry{-1, byte(i)}
		lookup[[2]int{-1, i}] = i
	}
	return dict, lookup
}
