package jsonfs

import (
	"fmt"
	"io/fs"
	"strings"
)

type file struct {
	path   string
	reader *strings.Reader
	value  interface{}
}

func (f *file) Stat() (fs.FileInfo, error) {
	return fs.FileInfo(&fileInfo{
		isDir: isDir(f.value),
		path:  f.path}), nil
}

func (f *file) Read(b []byte) (n int, err error) {
	if f.reader == nil {
		f.reader = strings.NewReader(fmt.Sprint(f.value))
	}
	return f.reader.Read(b)
}

func (f *file) Close() error {
	f.reader = nil
	return nil
}
