package service

import (
	"api/ent"
	"api/ent/file"
	"api/ent/mirror"
	"api/models"
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type SearchServiceImpl struct {
	dbClient    *ent.Client
	contentPath string
}

func NewSearchService(dbClient *ent.Client, contentPath string) *SearchServiceImpl {
	return &SearchServiceImpl{dbClient: dbClient, contentPath: contentPath}
}

func (s *SearchServiceImpl) ContentPath() string {
	return s.contentPath
}

func (s *SearchServiceImpl) ReadFileContents(fileName string) (string, error) {
	contentBytes, err := os.ReadFile(s.contentPath + fileName)
	if err != nil {
		return "", nil
	}

	return string(contentBytes), nil
}

func (s *SearchServiceImpl) GetFileDto(ctx context.Context, fileName string) (models.FileDTO, error) {
	fileMeta, err := s.dbClient.File.Query().Where(file.Path(fileName)).Only(ctx)
	if err != nil {
		return models.FileDTO{}, err
	}

	dto := models.FileDTO{
		Path:    fileMeta.Path,
		Title:   fileMeta.Title,
		Updated: fileMeta.Updated,
		Content: fileMeta.Content,
	}

	return dto, nil
}

func (s *SearchServiceImpl) UpdateIndex(ctx context.Context, filePath string) error {
	content, err := s.ReadFileContents(filePath)
	if err != nil {
		return err
	}
	if len(content) == 0 {
		return errors.New("file is empty")
	}

	article, meta := getArticle(content)

	if meta.MirrorOf != "" {
		return s.setMirror(filePath, meta.MirrorOf, ctx)
	}

	if meta.Title == "" {
		meta.Title, err = getFileTitle(article)
		if err != nil {
			return err
		}
	}

	updated, err := s.getLastModifiedTime(filePath)
	if err != nil {
		return err
	}

	return s.setFile(filePath, meta.Title, updated, article, ctx)
}

func getArticle(fileContent string) (article string, meta models.FileMetaDTO) {
	firstLine := fileContent
	if split := strings.Split(fileContent, "\n"); len(split) != 0 {
		firstLine = split[0]
	}

	if err := json.Unmarshal([]byte(firstLine), &meta); err == nil {
		article = fileContent[len(firstLine)-1 : len(fileContent)-1]
	} else {
		article = fileContent
		meta = models.FileMetaDTO{}
	}

	return article, meta
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

func (s *SearchServiceImpl) getLastModifiedTime(filePath string) (time.Time, error) {
	cmd := exec.Command("git", "log", "-1", "--format=%ct", filePath)
	cmd.Dir = s.contentPath

	output, err := cmd.Output()
	if err != nil {
		return time.Time{}, err
	}

	timestampStr := strings.Trim(string(output), "\n")
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

func (s *SearchServiceImpl) setFile(path string, title string, updated time.Time, content string, ctx context.Context) error {
	existingId, err := s.dbClient.File.Query().Where(file.Path(path)).OnlyID(ctx)

	if err != nil {
		_, err = s.dbClient.File.Update().
			Where(file.ID(existingId)).
			SetTitle(title).
			SetContent(content).
			SetUpdated(updated).
			Save(ctx)
		return err
	}

	_, err = s.dbClient.File.Create().
		SetPath(path).
		SetTitle(title).
		SetContent(content).
		SetUpdated(updated).
		Save(ctx)

	return err
}
