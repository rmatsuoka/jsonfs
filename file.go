package jsonfs

import (
	"fmt"
	"io/fs"
	"strings"
)

type File struct {
	name   string
	value  interface{}
	reader *strings.Reader
	rootFS *FS
	path   string
}

func (f *File) Stat() (fs.FileInfo, error) {
	// check existance
	if _, _, err := f.rootFS.namev(f.path); err != nil {
		return nil, err
	}
	return fs.FileInfo(&FileInfo{name: f.name, isDir: isDir(f.value)}), nil
}

func (f *File) Read(b []byte) (n int, err error) {
	if f.reader == nil {
		f.reader = strings.NewReader(fmt.Sprint(f.value))
	}
	return f.reader.Read(b)
}

func (f *File) Close() error {
	f.reader = nil
	return nil
}
