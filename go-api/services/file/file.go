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

func WriteFile(path string, data []byte) (ok bool) {
	err := os.WriteFile(path, data, 0755)

	return err == nil
}
