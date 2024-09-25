package models

type FileMetaDTO struct {
	MirrorOf string `json:"mirrorOf"`
	Title    string `json:"title"`
	Author   string `json:"author"`
}
