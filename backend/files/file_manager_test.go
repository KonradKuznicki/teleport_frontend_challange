package files_test

import (
	"testing"

	"challenge/files"
	"github.com/stretchr/testify/suite"
)

func TestFileManager(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestFileManagerSuite))
}

type TestFileManagerSuite struct {
	suite.Suite
	FileManager *files.FileManager
}

func (s *TestFileManagerSuite) SetupSuite() {
	s.FileManager, _ = files.NewFileManager("../../resources/traversable")
}

func (s *TestFileManagerSuite) TestFileManager() {
	s.NotNil(s.FileManager)
}

func (s *TestFileManagerSuite) TestFileManager_listFolderNonExistent() {
	_, err := s.FileManager.ListFolder("asdf")
	s.NotNil(err)
	s.Equal(err.Code(), files.FileManagerNotFound)
}

func (s *TestFileManagerSuite) TestFileManager_listEmpty() {
	folder, err := s.FileManager.ListFolder("empty")
	s.Nil(err)
	s.Equal([]*files.File{}, folder)
}

func (s *TestFileManagerSuite) TestFileManager_errorOnEscapism() {
	folder, err := s.FileManager.ListFolder("../../empty")
	s.NotNil(err)
	s.Equal(err.Code(), files.FileManagerInvalidPath)
	s.Equal([]*files.File(nil), folder)
}

func (s *TestFileManagerSuite) TestFileManager_directMatch() {
	folder, err := s.FileManager.ListFolder("notes.txt")
	s.Nil(err)
	s.Equal([]*files.File{
		{Name: "notes.txt", Type: "txt", Size: 15},
	}, folder)
}

func (s *TestFileManagerSuite) TestFileManager_list() {
	folder, err := s.FileManager.ListFolder("")
	s.Nil(err)
	s.Equal([]*files.File{
		{Name: "documents", Type: "folder", Size: 0},
		{Name: "empty", Type: "folder", Size: 0},
		{Name: "images", Type: "folder", Size: 0},
		{Name: "notes.txt", Type: "txt", Size: 15},
	}, folder)
}

func (s *TestFileManagerSuite) TestFileManager_listPathAlternatives() {
	folder, _ := s.FileManager.ListFolder("")
	folder1, _ := s.FileManager.ListFolder("/")
	s.Equal(folder, folder1)
	folder2, _ := s.FileManager.ListFolder("./")
	s.Equal(folder, folder2)
	folder3, _ := s.FileManager.ListFolder(".")
	s.Equal(folder, folder3)
	folder4, _ := s.FileManager.ListFolder("////")
	s.Equal(folder, folder4)
}
