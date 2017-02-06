package decoder

// Fuzz is used for fuzzer testing, ignore it.
func Fuzz(data []byte) int {
	err := Valid(data)
	if err != nil {
		return 0
	}

	return 1
}
