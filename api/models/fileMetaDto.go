package models

type FileMetaDto struct {
	MirrorOf string `json:"mirrorOf,omitempty"`
	Title    string `json:"title"`
	Path     string `json:"path"`
}
