package models

import "time"

type FileDTO struct {
	MirrorOf      string    `json:"mirrorOf,omitempty"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Path          string    `json:"path"`
	Updated       time.Time `json:"updated"`
	LastCommit    string    `json:"lastCommit"`
	Content       string    `json:"content,omitempty"`
	Headings      []Heading `json:"headings,omitempty"`
	IncomingLinks []Link    `json:"links,omitempty"`
}
