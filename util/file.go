package util

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(path string) []string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
