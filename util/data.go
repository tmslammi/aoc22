package util

import (
	"github.com/emirpasic/gods/utils"
	"sort"
)

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
	for _, row := range slice {
		if row < min {
			min = row
		}
	}
	return min
}

func FindMax(slice []int) int {
	max := 0
	for _, row := range slice {
		if row > max {
			max = row
		}
	}
	return max
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

func ChunkByChar(slice []string, char string) [][]string {
	var results [][]string
	var temp []string
	for i, row := range slice {
		if row == char {
			results = append(results, temp)
			temp = nil
		} else {
			temp = append(temp, row)
			if i == len(slice)-1 {
				results = append(results, temp)
			}
		}
	}
	return results
}

func Contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func InterfaceToSlice(rows []interface{}) []string {
	var results []string
	for _, row := range rows {
		results = append(results, utils.ToString(row))
	}
	return results
}
