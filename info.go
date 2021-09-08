package jsonfs

import (
	"io/fs"
	pathpkg "path"
	"time"
)

type fileInfo struct {
	isDir bool
	path  string
	size  int64
}

func (info *fileInfo) Name() string {
	return pathpkg.Base(info.path)
}

func (info *fileInfo) Size() int64 {
	return info.size
}

func (info *fileInfo) Mode() fs.FileMode {
	const rdonly = 0444
	const exonly = 0111
	if info.isDir {
		return rdonly | exonly | fs.ModeDir
	}
	return rdonly
}

func (info *fileInfo) Type() fs.FileMode {
	if info.isDir {
		return fs.ModeDir
	}
	return 0
}

func (info *fileInfo) ModTime() time.Time {
	return time.Time{}
}

func (info *fileInfo) IsDir() bool {
	return info.isDir
}

func (info *fileInfo) Info() (fs.FileInfo, error) {
	return fs.FileInfo(info), nil
}

func (info *fileInfo) Sys() interface{} {
	return nil
}
