package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator) // 路径分隔符，其实就是个分号
// Entry 是类路径的抽象
type Entry interface {
	String() string
	readClass(className string) ([]byte, Entry, error)
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {

	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
