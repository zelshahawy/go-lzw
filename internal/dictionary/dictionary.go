package dictionary

type DictionaryEntry struct {
	Prefix int  // index of the previous dictionary code
	Ch     byte // character appended to the prefix
}

func InitDictionary() ([]DictionaryEntry, map[[2]int]int) {
	dict := make([]DictionaryEntry, 256)
	lookup := make(map[[2]int]int, 256)
	// Initialize single-byte entries
	for i := 0; i < 256; i++ {
		dict[i] = DictionaryEntry{-1, byte(i)}
		lookup[[2]int{-1, i}] = i
	}
	return dict, lookup
}
