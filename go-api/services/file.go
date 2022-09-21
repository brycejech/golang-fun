package services

import (
	"fmt"
	"os"
)

type FileService interface {
	GetFileContents(path string) []byte
	WriteFile(path string, data []byte) bool
}

type fileService struct{}

func NewFileService() FileService {
	return &fileService{}
}

func (service *fileService) GetFileContents(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}

	return bytes
}

func (service *fileService) WriteFile(path string, data []byte) (ok bool) {
	err := os.WriteFile(path, data, 0755)

	return err == nil
}
