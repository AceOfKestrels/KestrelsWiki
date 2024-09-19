package service

import (
	"api/ent"
	"api/ent/file"
	"api/ent/mirror"
	"api/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
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
		return "", err
	}

	return string(contentBytes), nil
}
func (s *SearchServiceImpl) GetFileDto(context context.Context, fileName string) (models.FileDTO, error) {
	fileName = strings.ToLower(fileName)

	mirrorData, err := s.dbClient.Mirror.Query().Where(mirror.OriginPath(fileName)).Only(context)
	if err == nil {
		return models.FileDTO{MirrorOf: mirrorData.TargetPath}, nil
	}

	fileMeta, err := s.dbClient.File.Query().Where(file.Path(fileName)).Only(context)
	if err != nil {
		return models.FileDTO{}, err
	}

	dto := models.FileDTO{
		Path:    strings.ToLower(fileMeta.Path),
		Title:   fileMeta.Title,
		Updated: fileMeta.Updated,
		Content: fileMeta.Content,
	}

	return dto, nil
}

func (s *SearchServiceImpl) UpdateIndex(context context.Context) error {
	var directories []string
	directories = append(directories, "")

	var errs error

	for len(directories) > 0 {
		currentDir := directories[0]
		directories[0] = directories[len(directories)-1]
		directories = directories[:len(directories)-1]

		dirEntries, err := os.ReadDir(s.contentPath + "/" + currentDir)
		if err != nil {
			errors.Join(errs, err)
			continue
		}

		files, dirs := s.sortFilesAndDirs(currentDir, dirEntries)

		for _, d := range dirs {
			directories = append(directories, d)
		}

		for _, f := range files {
			errors.Join(errs, s.AddFileToIndex(context, f))
		}
	}

	return errs
}

func (s *SearchServiceImpl) sortFilesAndDirs(currentPath string, entries []os.DirEntry) (filePaths []string, dirPaths []string) {
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		fullPath := currentPath + "/" + entry.Name()
		if strings.HasPrefix(fullPath, "/") {
			fullPath = fullPath[1:]
		}

		if entry.IsDir() {
			dirPaths = append(dirPaths, fullPath)
		} else if strings.HasSuffix(entry.Name(), ".md") {
			filePaths = append(filePaths, fullPath)
		}
	}
	return filePaths, dirPaths
}

func (s *SearchServiceImpl) AddFileToIndex(context context.Context, filePath string) error {
	content, err := s.ReadFileContents(filePath)
	if err != nil {
		return s.getFileReadingError(filePath, err.Error())
	}
	if len(content) == 0 {
		return s.getFileReadingError(filePath, "file is empty")
	}

	article, meta := getArticle(content)

	if meta.MirrorOf != "" {
		return s.setMirror(filePath, meta.MirrorOf, context)
	}

	if meta.Title == "" {
		meta.Title, err = getFileTitle(article)
		if err != nil {
			return s.getFileReadingError(filePath, err.Error())
		}
	}

	updated, err := s.getLastModifiedTime(filePath)
	if err != nil {
		return s.getFileReadingError(filePath, err.Error())
	}

	return s.setFile(filePath, meta.Title, updated, article, context)
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
	origin = strings.ToLower(origin)

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
	path = strings.ToLower(path)

	existingId, err := s.dbClient.File.Query().Where(file.Path(path)).OnlyID(ctx)

	if err == nil {
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

func (s *SearchServiceImpl) getFileReadingError(fileName string, errorMessage string) error {
	return fmt.Errorf("error reading file at %v: %v", fileName, errorMessage)
}

func (s *SearchServiceImpl) SearchFiles(ctx context.Context, search models.SearchContext) ([]string, error) {
	found := mapset.NewSet[string]()

	var errs error

	foundInTitles, err := s.dbClient.File.
		Query().
		Where(
			file.And(
				file.Not(
					file.Path(search.CurrentPage)),
				file.TitleContainsFold(search.SearchString))).
		All(ctx)
	if err != nil {
		errors.Join(errs, err)
	}
	for _, f := range foundInTitles {
		found.Add(f.Path)
	}

	foundInArticle, err := s.dbClient.File.
		Query().
		Where(
			file.And(
				file.Not(
					file.Path(search.CurrentPage)),
				file.ContentContainsFold(search.SearchString))).
		All(ctx)
	if err != nil {
		errors.Join(errs, err)
	}
	for _, f := range foundInArticle {
		found.Add(f.Path)
	}

	return found.ToSlice(), errs
}
