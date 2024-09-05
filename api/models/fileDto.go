package models

import "time"

type FileDTO struct {
	MirrorOf      string    `json:"mirrorOf,omitempty"`
	Title         string    `json:"title"`
	Path          string    `json:"path"`
	Updated       time.Time `json:"updated"`
	Content       string    `json:"content"`
	Headings      []Heading `json:"headings,omitempty"`
	IncomingLinks []Link    `json:"links,omitempty"`
}
