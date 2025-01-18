package dictionary

// DictionaryEntry is a struct that holds the prefix and the character appended to it.
// it has the following fields: Prefix and Ch.
type DictionaryEntry struct {
	Prefix int  // index of the previous dictionary code
	Ch     byte // character appended to the prefix
}

// InitDictionary initializes the dictionary with single-byte entries.
// It returns a slice of DictionaryEntry and a map of [2]int to int.
// The first 256 entries are initialized with single-byte entries.
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
