package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	FileManagerNotFound = iota
	FileManagerInvalidPath
	FileManagerReadError
)

type FileManager struct {
	rootAbs string
}

type File struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Type string `json:"type"`
}

type FileManagerError interface {
	error
	Code() int
}

type fmerr struct {
	err  error
	code int
}

func (f *fmerr) Error() string {
	return f.err.Error()
}

func (f *fmerr) Code() int {
	return f.code
}

func (m *FileManager) ListFolder(subPath string) ([]*File, FileManagerError) {
	path, err := m.computePath(subPath)
	if err != nil {
		return nil, err
	}

	listing, err := m.tryExactFileMatch(path)
	if err != nil || listing != nil {
		return listing, err
	}

	return m.readFiles(path)
}

func (m *FileManager) readFiles(path string) ([]*File, FileManagerError) {
	listing := []*File{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, &fmerr{err: err, code: FileManagerNotFound}
		}
		return nil, &fmerr{err: err, code: FileManagerReadError}
	}

	for _, file := range files {
		ftype, size := itemType(file)
		listing = append(listing, &File{Name: file.Name(), Size: size, Type: ftype})
	}
	return listing, nil
}

func (m *FileManager) tryExactFileMatch(path string) ([]*File, FileManagerError) {
	fi, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, &fmerr{err: err, code: FileManagerNotFound}
		}
		return nil, &fmerr{err: err, code: FileManagerReadError}
	}

	if fi.IsDir() {
		return nil, nil
	}

	ftype := computeFileType(fi.Name())

	return []*File{{
		Name: fi.Name(),
		Size: fi.Size(),
		Type: ftype,
	}}, nil
}

func computeFileType(name string) string {
	ftype := "unknown"
	split := strings.Split(name, ".")
	if len(split) > 0 {
		ftype = split[len(split)-1]
	}
	return ftype
}

func (m *FileManager) computePath(subPath string) (string, FileManagerError) {
	path, err := filepath.Abs(filepath.Join(m.rootAbs, subPath))
	if err != nil {
		return "", &fmerr{err: err, code: FileManagerInvalidPath}
	}

	if strings.Index(path, m.rootAbs) != 0 {
		return "", &fmerr{err: fmt.Errorf("attempt at escaping path: %s", subPath), code: FileManagerInvalidPath}
	}
	return path, nil
}

func (m *FileManager) FilesHandler(writer http.ResponseWriter, request *http.Request) {
	apiPrefix := "/API/v1/files"
	if request.URL.Path[:len(apiPrefix)] != apiPrefix {
		log.Println("bad request", request.URL.Path)
		http.Error(writer, "bad request", http.StatusBadRequest)
		return
	}
	log.Println("trying path ", request.URL.Path)

	path := request.URL.Path[len(apiPrefix):]

	log.Println("checking path: ", path)
	list, ferr := m.ListFolder(path)
	if ferr != nil {
		switch ferr.Code() {
		case FileManagerInvalidPath:
			log.Println("path invalid: ", path)
			http.Error(writer, "path not found", http.StatusNotFound)
			return
		case FileManagerNotFound:
			log.Println("path not found: ", path)
			http.Error(writer, "path not found", http.StatusNotFound)
			return

		case FileManagerReadError:
			log.Println("path read error: ", ferr.Error())
			http.Error(writer, "server error", http.StatusInternalServerError)
			return
		default:
			log.Println("unknown bad things: ", ferr.Error())
			http.Error(writer, "server error", http.StatusInternalServerError)
		}
	}

	marshal, err := json.Marshal(list)
	if err != nil {
		log.Printf("error marshaling response %v", err)
		http.Error(writer, "server error", http.StatusInternalServerError)
		return
	}

	writer.Header().Add("content-type", "application/json")
	_, err = io.WriteString(writer, string(marshal))
	if err != nil {
		log.Printf("error writeing response for %s error: %v", request.URL.Host, err)
		http.Error(writer, "server error", http.StatusInternalServerError)
	}
	log.Println("response should be ok")
}

func itemType(file fs.FileInfo) (string, int64) {
	if file.IsDir() {
		return "folder", 0
	}
	return computeFileType(file.Name()), file.Size()
}

func NewFileManager(rootPath string) (*FileManager, error) {
	abs, err := filepath.Abs(rootPath)
	if err != nil {
		return nil, err
	}

	return &FileManager{
		rootAbs: abs,
	}, nil
}
