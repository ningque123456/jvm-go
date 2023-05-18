package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	abs string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{abs: absPath}
}
func (self *ZipEntry) String() string {
	return self.abs
}

// 解析zip文件中的每个file
// 每次调用都要Open -> Close 整个zip文件，性能较低。
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.abs)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			file, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer file.Close()
			data, err := io.ReadAll(file)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
