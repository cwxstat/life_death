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

func DataFile(file string) string {
	if filepath.IsAbs(file) {
		return file
	}
	return filepath.Join(basepath, file)
}
