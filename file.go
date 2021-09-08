package jsonfs

import (
	"fmt"
	"io/fs"
	"strings"
)

type File struct {
	path   string
	reader *strings.Reader
	value  interface{}
}

func (f *File) Stat() (fs.FileInfo, error) {
	return fs.FileInfo(&FileInfo{
		isDir: isDir(f.value),
		path:  f.path}), nil
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
