package util

import (
	"log"
	"os"
)

func DoesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func WriteFile(fileName string, content string) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}
