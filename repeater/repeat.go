package iterations

// Repeat takes one letter and returns it 5 times
func Repeat(letter string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += letter
	}
	return repeated
}
