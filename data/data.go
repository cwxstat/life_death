package data

import (
	"path/filepath"
	"runtime"
)

// basepath is the root directory of this package.
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

func DataFile() string {
	file := "lifeexp.txt"
	if filepath.IsAbs(file) {
		return file
	}

	return filepath.Join(basepath, file)
}
