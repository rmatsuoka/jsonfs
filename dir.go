package jsonfs

import (
	"io/fs"
)

type DirEntry struct {
	info   *FileInfo
	rootFS *FS
	path   string
}

func (d *DirEntry) Name() string {
	return d.info.Name()
}

func (d *DirEntry) IsDir() bool {
	return d.info.IsDir()
}

func (d *DirEntry) Type() fs.FileMode {
	return d.info.Mode()
}

func (d *DirEntry) Info() (fs.FileInfo, error) {
	return fs.Stat(d.rootFS, d.path)
}
