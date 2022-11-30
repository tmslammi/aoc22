package util

import (
	"log"
	"strconv"
)

func ToInt(s string) int {
	parsed, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}

func ToString(i int) string {
	return strconv.Itoa(i)
}
