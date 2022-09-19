package fileService

import (
	"fmt"
	"os"
)

func GetFileContents(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}

	return bytes
}
