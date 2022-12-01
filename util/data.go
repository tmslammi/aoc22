package util

import "sort"

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

func FindMin(slice []int) int {
	min := 0
	prev := 0
	for _, row := range slice {
		if row < prev {
			min = prev
		}
		prev = row
	}
	return min
}

func SortDesc(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

func SortAsc(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice
}

func Sum(slice []int) int {
	sum := 0
	for _, row := range slice {
		sum += row
	}
	return sum
}
