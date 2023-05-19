package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	//cp := &Classpath{}
	cp := new(Classpath)
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLib := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = newWildCardEntry(jreLib)
	// jre/lib/ext/*
	jreExt := filepath.Join(jreLib, "ext", "*")
	cp.extClasspath = newWildCardEntry(jreExt)
}

// 先找Xjre参数 ， 然后找当前目录 ， 最后找环境变量
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder !")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 先找cp参数 ， 否则默认为当前目录
func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = newEntry(cpOption)
}
func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := cp.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	return cp.userClasspath.readClass(className)
}
func (cp *Classpath) String() string {
	return cp.userClasspath.String()
}
