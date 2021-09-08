package jsonfs

import (
	"io/fs"
	pathpkg "path"
	"time"
)

type FileInfo struct {
	isDir bool
	path  string
	size  int64
}

func (info *FileInfo) Name() string {
	return pathpkg.Base(info.path)
}

func (info *FileInfo) Size() int64 {
	return info.size
}

func (info *FileInfo) Mode() fs.FileMode {
	const rdonly = 0444
	const exonly = 0111
	if info.isDir {
		return rdonly | exonly | fs.ModeDir
	}
	return rdonly
}

func (info *FileInfo) Type() fs.FileMode {
	if info.isDir {
		return fs.ModeDir
	}
	return 0
}

func (info *FileInfo) ModTime() time.Time {
	return time.Time{}
}

func (info *FileInfo) IsDir() bool {
	return info.isDir
}

func (info *FileInfo) Info() (fs.FileInfo, error) {
	return fs.FileInfo(info), nil
}

func (info *FileInfo) Sys() interface{} {
	return nil
}
