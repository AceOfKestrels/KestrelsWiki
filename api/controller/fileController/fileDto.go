package fileController

import (
	"errors"
	"strings"
)

type FileDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewFileDto(content string) (*FileDTO, error) {
	if len(content) == 0 {
		return nil, errors.New("content is required")
	}

	title, error := getFileTitle(content)
	if error != nil || len(title) == 0 {
		return nil, errors.New("title is required")
	}

	return &FileDTO{title, content}, nil
}

func getFileTitle(content string) (string, error) {
	split := strings.Split(content, "\n")
	firstLine := content
	if len(split) != 0 {
		firstLine = split[0]
	}
	if !strings.HasPrefix(firstLine, "# ") {
		return "", errors.New("content contains no title")
	}
	return strings.Trim(firstLine[2:len(firstLine)-1], " "), nil
}
