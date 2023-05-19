package classpath

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func newWildCardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove "*"
	var compositeEntry []Entry
	filepath.Walk(baseDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	})
	return compositeEntry
}
