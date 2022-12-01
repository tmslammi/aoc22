package util

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Read(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatal(err)
	}
	str := string(file)
	return str
}

func ReadLines(path string) []string {
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
