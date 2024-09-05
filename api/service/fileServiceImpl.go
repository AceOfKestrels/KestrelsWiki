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

	split := strings.Split(fileContent, "\n")
	firstLine := fileContent
	if len(split) != 0 {
		firstLine = split[0]
	}

	var dto models.FileDTO
	if err = json.Unmarshal([]byte(firstLine), &dto); err == nil {
		dto.Content = fileContent[len(firstLine)-1 : len(fileContent)-1]
	} else {
		dto = models.FileDTO{Content: fileContent}
	}

	if dto.MirrorOf != "" {
		return models.FileDTO{MirrorOf: dto.MirrorOf}, nil
	}

	if dto.Title == "" {
		dto.Title, err = getFileTitle(dto.Content)
		if err != nil {
			return models.FileDTO{}, err
		}
	}

	return dto, nil
}

func getFileTitle(content string) (string, error) {
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		lineTrimmed := strings.TrimSpace(line)
		if len(lineTrimmed) == 0 {
			continue
		}

		if strings.HasPrefix(line, "# ") {
			return strings.Trim(line[2:len(line)-1], " "), nil
		}

		break
	}

	return "", errors.New("no title found")
}
