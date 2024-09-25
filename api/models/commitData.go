package models

import (
	"encoding/json"
	"time"
)

type CommitData struct {
	Hash   string    `json:"hash"`
	Date   time.Time `json:"date"`
	Author string    `json:"author"`
}

func ParseCommitData(jsonData []byte) (commitData CommitData, err error) {
	err = json.Unmarshal(jsonData, &commitData)
	return
}
