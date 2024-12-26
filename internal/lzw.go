package internal

func EncodeFile(fileName string) error {
	return execEncoding(fileName)
}

func DecodeFile(fileName string) error {
	return execDecoding(fileName)
}
