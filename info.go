package jsonfs

import (
	"io/fs"
	"time"
)

type FileInfo struct {
	name  string
	size  int64
	isDir bool
}

func (info *FileInfo) Name() string {
	return info.name
}

func (info *FileInfo) Size() int64 {
	return info.size
}

func (info *FileInfo) Mode() fs.FileMode {
	const rdonly = 0444
	if info.isDir {
		return rdonly | fs.ModeDir
	}
	return rdonly
}

func (info *FileInfo) ModTime() time.Time {
	return *new(time.Time)
}

func (info *FileInfo) IsDir() bool {
	return info.isDir
}

func (info *FileInfo) Sys() interface{} {
	return nil
}
