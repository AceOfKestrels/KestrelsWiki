package service

import (
	"api/ent"
	"api/ent/file"
	"api/ent/mirror"
	"api/models"
	"context"
	"encoding/json"
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type FileService interface {
	ReadFileContents(fileName string) (string, error)
}

type SearchServiceImpl struct {
	fileService FileService
	dbClient    *ent.Client
}

func NewSearchService(fileService FileService, dbClient *ent.Client) *SearchServiceImpl {
	return &SearchServiceImpl{fileService: fileService, dbClient: dbClient}
}

func (s *SearchServiceImpl) UpdateIndex(ctx context.Context, filePath string) error {
	content, err := s.fileService.ReadFileContents(filePath)
	if err != nil {
		return err
	}
	if len(content) == 0 {
		return errors.New("file is empty")
	}

	firstLine := content
	if split := strings.Split(content, "\n"); len(split) != 0 {
		firstLine = split[0]
	}

	var meta models.FileMetaDTO
	var articleContent string
	if err = json.Unmarshal([]byte(firstLine), &meta); err == nil {
		articleContent = content[len(firstLine)-1 : len(content)-1]
	} else {
		articleContent = content
		meta = models.FileMetaDTO{}
	}

	if meta.MirrorOf != "" {
		return s.setMirror(filePath, meta.MirrorOf, ctx)
	}

	if meta.Title == "" {
		meta.Title, err = getFileTitle(articleContent)
		if err != nil {
			return err
		}
	}

	updated, err := getLastModifiedTime(filePath)
	if err != nil {
		return err
	}

	existingId, err := s.dbClient.File.Query().Where(file.Path(filePath)).OnlyID(ctx)

	if err != nil {
		_, err = s.dbClient.File.Update().
			Where(file.ID(existingId)).
			SetTitle(meta.Title).
			SetContent(articleContent).
			SetUpdated(updated).
			Save(ctx)
		return err
	}

	_, err = s.dbClient.File.Create().
		SetPath(filePath).
		SetTitle(meta.Title).
		SetContent(articleContent).
		SetUpdated(updated).
		Save(ctx)

	return err
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

func getLastModifiedTime(filePath string) (time.Time, error) {
	cmd := exec.Command("git", "log", "-1", "--format=%ct", filePath)

	output, err := cmd.Output()
	if err != nil {
		return time.Time{}, err
	}

	timestampStr := string(output)
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(timestamp, 0), nil
}

func (s *SearchServiceImpl) setMirror(origin string, target string, ctx context.Context) error {
	id, err := s.dbClient.Mirror.Query().Where(mirror.OriginPath(origin)).OnlyID(ctx)
	if err == nil {
		_, err = s.dbClient.Mirror.Update().Where(mirror.ID(id)).SetTargetPath(target).Save(ctx)
		if err != nil {
			return err
		}
	}

	_, err = s.dbClient.Mirror.Create().SetOriginPath(origin).SetTargetPath(target).Save(ctx)
	return err
}
