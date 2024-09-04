package service

import (
	"os"
)

type FileServiceImpl struct{}

func NewFileService() *FileServiceImpl {
	return &FileServiceImpl{}
}

func (f *FileServiceImpl) ReadFileContents(fileName string) (string, error) {
	contentBytes, error := os.ReadFile(fileName)
	if error != nil {
		return "", nil
	}

	return string(contentBytes), nil
}
