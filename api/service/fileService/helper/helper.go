// Package helper
// contains functions to assist in reading files and parsing article content
package helper

import (
	"api/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"os"
	"os/exec"
	"strings"
)

// GetFilesAndDirs
// searches through each directory in <entries> and returns the full paths
// of any markdown files and subdirectories separately.
func GetFilesAndDirs(currentPath string, entries []os.DirEntry) (filePaths []string, dirPaths []string) {
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

// ReadFileContents
// reads the contents of a file, returning its contents as a string.
// If any error occurs, return an empty string as well as the error.
func ReadFileContents(path string) (content string, err error) {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(contentBytes), nil
}

// GetArticle
// given the entire content of an article file will return the article body
// and the deserialized article meta.
func GetArticle(fileContent string) (article string, meta models.FileMetaDTO) {
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

// GetArticleTitle
// Given the content of an article will find the first H1 and return its contents.
// The function will return an empty string and an error if no H1 is found.
func GetArticleTitle(content string) (string, error) {
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

// GetCommitData
// given a file path relative to a Git repository's root and the path of a Git repository
// use git log to retrieve information about the last commit the file was edited in.
// If any errors occur, a zero commit data and the error are returned.
func GetCommitData(filePath string, repositoryPath string) (commitData models.CommitData, err error) {
	output, err := ExecuteCommand(repositoryPath, "git", "log", "-1",
		`--pretty=format:{"hash": "%H", "date": "%ad", "author": "%an"}`, "--date=iso-strict", filePath)
	if err != nil {
		return
	}

	commitData, err = models.ParseCommitData(output)
	return
}

// GetFileReadingError
// returns an error with a message formatted like:
// "error reading file at fileName: errorMessage"
func GetFileReadingError(fileName string, errorMessage string) error {
	return fmt.Errorf("error reading file at %v: %v", fileName, errorMessage)
}

func RenderMarkdown(md string) (renderedHtml string) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}

func ExecuteCommand(directory string, command string, args ...string) (output []byte, err error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = directory
	output, err = cmd.Output()
	return
}
