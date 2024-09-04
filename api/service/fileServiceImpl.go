package service

import (
	"api/models"
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type FileServiceImpl struct{}

func NewFileService() *FileServiceImpl {
	return &FileServiceImpl{}
}

func (f *FileServiceImpl) ReadFileContents(fileName string) (string, error) {
	contentBytes, err := os.ReadFile(fileName)
	if err != nil {
		return "", nil
	}

	return string(contentBytes), nil
}

func (f *FileServiceImpl) GetFileDto(fileName string) (models.FileDTO, error) {
	fileContent, err := f.ReadFileContents(fileName)
	if err != nil {
		return models.FileDTO{}, err
	}

	if len(fileContent) == 0 {
		return models.FileDTO{}, errors.New("content is required")
	}

	meta, remainingContent := getFileMeta(fileContent)

	if meta.MirrorOf != "" {
		return models.FileDTO{Meta: meta}, nil
	}

	if meta.Title == "" {
		title, err := getFileTitle(remainingContent)
		if err != nil {
			return models.FileDTO{}, err
		}
		meta.Title = title
	}

	return models.FileDTO{Content: remainingContent, Meta: meta}, nil
}

func getFileTitle(content string) (string, error) {
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if len(strings.TrimSpace(line)) != 0 {
			if strings.HasPrefix(line, "# ") {
				return strings.Trim(line[2:len(line)-1], " "), nil
			}
			break
		}
	}

	return "", nil
}

func getFileMeta(content string) (models.FileMetaDto, string) {
	split := strings.Split(content, "\n")
	firstLine := content
	if len(split) != 0 {
		firstLine = split[0]
	}

	var dto models.FileMetaDto
	if err := json.Unmarshal([]byte(firstLine), &dto); err != nil {
		return models.FileMetaDto{}, content
	}

	return dto, content[len(firstLine)-1 : len(content)-1]
}
