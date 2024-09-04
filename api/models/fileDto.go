package models

type FileDTO struct {
	Content string      `json:"content"`
	Meta    FileMetaDto `json:"meta"`
}
