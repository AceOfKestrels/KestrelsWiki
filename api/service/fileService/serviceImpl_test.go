package fileService

import "testing"

// New_returnsServiceImplWithContentPath calls New, wants ServiceImpl with same path
func TestNew_returnsServiceImplWithContentPath(t *testing.T) {
	path := "path"

	serviceImpl := New(path)

	if serviceImpl.ContentPath() != path {
		t.Fatalf(`Want "%v", got "%v"`, path, serviceImpl.ContentPath())
	}
}
