package classpath

import (
	"archive/zip"
	"errors"
	"io"
)

// ZipEntry2 将zip文件作为结构体属性存放结构体中，只需readClass时open一次即可
type ZipEntry2 struct {
	abs   string
	zipRC *zip.ReadCloser
}

func (self *ZipEntry2) readClass(className string) ([]byte, Entry, error) {
	if self.zipRC == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}
	class := self.findClass(className)
	if class == nil {
		return nil, nil, errors.New("class not found: " + className)
	}
	data, err := readClass(class)
	if err != nil {
		return nil, nil, err
	}
	return data, self, nil
}
func (self *ZipEntry2) openJar() error {
	r, err := zip.OpenReader(self.abs)
	if err == nil {
		self.zipRC = r
	}
	return err
}
func (self ZipEntry2) findClass(className string) *zip.File {
	for _, f := range self.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}
func readClass(classFile *zip.File) ([]byte, error) {
	open, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(open)
	open.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (self *ZipEntry2) String() string {
	return self.abs
}
