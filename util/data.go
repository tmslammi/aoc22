package util

func Chunk(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func SliceToInt(slice []string) []int {
	var results []int
	for i := 0; i < len(slice); i++ {
		results = append(results, ToInt(slice[i]))
	}
	return results
}
