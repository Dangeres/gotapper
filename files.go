package main

import (
	"os"
	"strings"
)

func writeJson(path string, data []byte) {
	splittedPath := strings.Split(path, "/")

	foldPath := strings.Join(splittedPath[0:len(splittedPath)-1], "")

	os.MkdirAll(foldPath, 0777)

	err := os.WriteFile(path, data, 0777)

	if err != nil {
		panic(err)
	}
}

func readJson(path string) []byte {
	result, _ := os.ReadFile(path)

	return result
}
