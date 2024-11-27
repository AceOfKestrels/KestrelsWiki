package fileService

import (
	"api/logger"
	"api/models"
	"api/service/fileService/helper"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

const AlreadyUpToDate = "Already up to date."

var countFiles int
var countMirrors int

var mirrors map[string]string
var files map[string]models.FileDTO

type ServiceImpl struct {
	contentPath string
}

// New returns a new instance of ServiceImpl.
func New(contentPath string) *ServiceImpl {
	return &ServiceImpl{contentPath: contentPath}
}

// ContentPath returns the path to the base directory of the content repository.
func (s *ServiceImpl) ContentPath() string {
	return s.contentPath
}

// GetFileDto searches the index for the data associated with a file located at filePath.
//
// filePath: the relative path from ContentPath
func (s *ServiceImpl) GetFileDto(filePath string) (models.FileDTO, error) {
	filePath = strings.ToLower(filePath)

	mirrorData, ok := mirrors[filePath]
	if ok {
		return models.FileDTO{MirrorOf: mirrorData}, nil
	}

	fileDto, ok := files[filePath]
	if !ok {
		return models.FileDTO{}, errors.New("file not found")
	}

	return fileDto, nil
}

func (s *ServiceImpl) Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func (s *ServiceImpl) GetArticle(path string) (string, error) {
	return "", nil
}

func (s *ServiceImpl) RebuildIndex(firstBuild bool) error {
	logger.Println(logger.INIT, "pulling changes from github")
	err := s.gitPull()
	if err != nil && !firstBuild {
		logger.Println(logger.INIT, "error pulling changes: "+err.Error())
		return err
	}

	logger.Println(logger.INIT, "rebuilding file index")
	err = s.updateIndex()

	if err != nil {
		logger.Println(logger.INIT, "error saving file index: "+err.Error())
	} else {
		logger.Println(logger.INIT, "finished building file index: %v files, %v mirrors", countFiles, countMirrors)
	}
	return err
}

func (s *ServiceImpl) gitPull() (err error) {
	output, err := helper.ExecuteCommand(s.contentPath, "git", "pull")
	if strings.Contains(string(output), AlreadyUpToDate) {
		err = fmt.Errorf("no changes available in remote repository")
	}

	return
}

// UpdateIndex recursively searches all subfolders for markdown files, starting at ContentPath,
// and adds them to the index.
func (s *ServiceImpl) updateIndex() error {
	var directories []string
	directories = append(directories, "")

	countFiles = 0
	countMirrors = 0
	mirrors = make(map[string]string)
	files = make(map[string]models.FileDTO)

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
			errors.Join(errs, s.AddFileToIndex(f))
		}
	}

	return errs
}

// AddFileToIndex reads the file located at filePath, adding its content and metadata to the index.
//
// filePath: the relative path from ContentPath
func (s *ServiceImpl) AddFileToIndex(filePath string) error {
	content, err := helper.ReadFileContents(s.contentPath + filePath)
	if err != nil {
		return helper.GetFileReadingError(filePath, err.Error())
	}
	if len(content) == 0 {
		return helper.GetFileReadingError(filePath, "file is empty")
	}

	article, meta := helper.GetArticle(content)

	if meta.MirrorOf != "" {
		s.setMirror(filePath, meta.MirrorOf)
		return nil
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

	renderedArticle := helper.RenderMarkdown(article)

	s.setFile(filePath, meta.Title, commitData, renderedArticle)
	return nil
}

// setMirror adds a mirror path to the mirrors map.
func (s *ServiceImpl) setMirror(origin string, target string) {
	origin = strings.ToLower(origin)
	mirrors[origin] = target
	countMirrors++
}

// setFile adds file content and metadata for a file located at path to the files map.
func (s *ServiceImpl) setFile(path string, title string, commitData models.CommitData, content string) {
	path = strings.ToLower(path)

	files[path] = models.FileDTO{
		Path:       path,
		Title:      title,
		Content:    content,
		Updated:    commitData.Date,
		Author:     commitData.Author,
		LastCommit: commitData.Hash,
	}

	countFiles++
}

// SearchFiles searches the indexed files and returns all the match search.
func (s *ServiceImpl) SearchFiles(search models.SearchContext) []models.FileDTO {
	search.SearchString = strings.ToLower(search.SearchString)
	search.CurrentPage = strings.ToLower(search.CurrentPage)

	found, result := s.searchFiles(map[string]bool{}, func(dto models.FileDTO) bool {
		return dto.Path != search.CurrentPage && strings.Contains(strings.ToLower(dto.Title), search.SearchString)
	})

	if !search.SearchInContent {
		return result
	}

	paths, files := s.searchFiles(map[string]bool{}, func(dto models.FileDTO) bool {
		return dto.Path != search.CurrentPage && strings.Contains(strings.ToLower(dto.Content), search.SearchString)
	})
	result = append(result, files...)
	found = s.addAllToMap(found, paths)

	return result
}

// searchFiles searches the files map using predicate and returns all results that are not in alreadyFound
func (s *ServiceImpl) searchFiles(alreadyFound map[string]bool, predicate func(dto models.FileDTO) bool) (pathsFound map[string]bool, dtos []models.FileDTO) {
	for _, dto := range files {
		if _, ok := alreadyFound[dto.Path]; predicate(dto) && !ok {
			dtos = append(dtos, dto)
			alreadyFound[dto.Path] = true
		}
	}

	slices.SortStableFunc(dtos, func(a, b models.FileDTO) int { return len(a.Title) - len(b.Title) })

	return pathsFound, dtos
}

func (s *ServiceImpl) addAllToMap(old map[string]bool, add map[string]bool) (result map[string]bool) {
	for key, value := range add {
		old[key] = value
	}
	return old
}
