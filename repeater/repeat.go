package iterations

// Repeat takes one letter and returns it 5 times
func Repeat(letter string, numOfTimes int) string {
	var repeated string
	for i := 0; i < numOfTimes; i++ {
		repeated += letter
	}
	return repeated
}
