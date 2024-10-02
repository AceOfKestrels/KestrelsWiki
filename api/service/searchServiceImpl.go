package service

import (
	"api/db/ent"
	"api/db/ent/file"
	"api/db/ent/mirror"
	"api/db/ent/predicate"
	"api/models"
	"api/service/helper"
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"os"
	"slices"
	"strings"
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

		files, dirs := helper.GetFilesAndDirs(currentDir, dirEntries)

		for _, d := range dirs {
			directories = append(directories, d)
		}

		for _, f := range files {
			errors.Join(errs, s.AddFileToIndex(context, f))
		}
	}

	return errs
}

func (s *SearchServiceImpl) AddFileToIndex(context context.Context, filePath string) error {
	content, err := helper.ReadFileContents(s.contentPath + filePath)
	if err != nil {
		return helper.GetFileReadingError(filePath, err.Error())
	}
	if len(content) == 0 {
		return helper.GetFileReadingError(filePath, "file is empty")
	}

	article, meta := helper.GetArticle(content)

	if meta.MirrorOf != "" {
		return s.setMirror(filePath, meta.MirrorOf, context)
	}

	if meta.Title == "" {
		meta.Title, err = helper.GetArticleTitle(article)
		if err != nil {
			return helper.GetFileReadingError(filePath, err.Error())
		}
	}

	commitData, err := helper.GetCommitData(filePath, s.contentPath)
	if err != nil {
		return helper.GetFileReadingError(filePath, err.Error())
	}

	return s.setFile(filePath, meta.Title, commitData, article, context)
}

func (s *SearchServiceImpl) setMirror(origin string, target string, ctx context.Context) (err error) {
	origin = strings.ToLower(origin)

	err = s.dbClient.Mirror.
		Create().
		SetOriginPath(origin).
		SetTargetPath(target).
		OnConflict(sql.ResolveWithNewValues()).
		Exec(ctx)
	return
}

func (s *SearchServiceImpl) setFile(path string, title string, commitData models.CommitData, content string, ctx context.Context) (err error) {
	path = strings.ToLower(path)

	err = s.dbClient.File.Create().
		SetPath(path).
		SetTitle(title).
		SetContent(content).
		SetUpdated(commitData.Date).
		SetAuthor(commitData.Author).
		SetCommitHash(commitData.Hash).
		OnConflict(sql.ResolveWithNewValues()).
		Exec(ctx)

	return
}

func (s *SearchServiceImpl) SearchFiles(ctx context.Context, search models.SearchContext) ([]models.FileDTO, error) {
	found, result, err := s.searchFiles(ctx, map[string]bool{},
		file.And(
			file.Not(
				file.Path(search.CurrentPage)),
			file.TitleContainsFold(search.SearchString)))
	if err != nil {
		return nil, err
	}

	if !search.SearchInContent {
		return result, nil
	}

	paths, files, err := s.searchFiles(ctx, found,
		file.And(
			file.Not(
				file.Path(search.CurrentPage)),
			file.ContentContainsFold(search.SearchString)))
	if err != nil {
		return nil, err
	}
	result = append(result, files...)
	found = s.addAllToMap(found, paths)

	return result, nil
}

func (s *SearchServiceImpl) searchFiles(ctx context.Context, alreadyFound map[string]bool, predicate predicate.File) (pathsFound map[string]bool, dtos []models.FileDTO, err error) {
	found, err := s.dbClient.File.Query().Where(predicate).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	for _, f := range found {
		if _, exists := alreadyFound[f.Path]; !exists {
			alreadyFound[f.Path] = true
			dtos = append(dtos, models.FileDTO{Path: f.Path, Title: f.Title, Updated: f.Updated})
		}
	}

	slices.SortStableFunc(dtos, func(a, b models.FileDTO) int { return len(a.Title) - len(b.Title) })

	return pathsFound, dtos, nil
}

func (s *SearchServiceImpl) addAllToMap(old map[string]bool, add map[string]bool) (result map[string]bool) {
	for key, value := range add {
		old[key] = value
	}
	return old
}
