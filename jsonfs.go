package jsonfs

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"sort"
	"strconv"
	"strings"
)

type FS struct {
	value interface{}
}

func NewFS(data []byte) (*FS, error) {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}
	return &FS{value}, nil
}

func (fsys *FS) namev(name string) (string, interface{}, error) {
	if !fs.ValidPath(name) {
		return "", nil, fs.ErrInvalid
	}

	if name == "." {
		return name, fsys.value, nil
	}

	var base string
	var value interface{} = fsys.value
	elems := strings.Split(name, "/")

	for len(elems) > 0 {
		base = elems[0]
		switch vv := value.(type) {
		case []interface{}:
			u, err := strconv.ParseUint(base, 10, 0)
			if int(u) >= len(vv) || err != nil {
				return "", nil, fs.ErrNotExist
			}
			value = vv[u]
		case map[string]interface{}:
			var ok bool
			if value, ok = vv[base]; !ok {
				return "", nil, fs.ErrNotExist
			}
		default:
			return "", nil, fs.ErrNotExist
		}
		elems = elems[1:]
	}
	return base, value, nil
}

func (fsys *FS) Open(name string) (fs.File, error) {
	base, value, err := fsys.namev(name)
	if err != nil {
		return nil, &fs.PathError{"open", name, err}
	}
	return fs.File(&File{
		name:   base,
		value:  value,
		reader: nil,
		rootFS: fsys,
		path:   name}), nil
}

func isDir(value interface{}) bool {
	switch value.(type) {
	case []interface{}, map[string]interface{}:
		return true
	}
	return false
}

func join(dir, base string) string {
	if dir == "." {
		return base
	}
	return dir + "/" + base
}

func (fsys *FS) ReadDir(name string) ([]fs.DirEntry, error) {
	_, value, err := fsys.namev(name)
	if err != nil {
		return nil, &fs.PathError{"readdir", name, err}
	}

	var dirs []fs.DirEntry
	switch vv := value.(type) {
	case []interface{}:
		dirs = make([]fs.DirEntry, len(vv))
		for i, v := range vv {
			base := strconv.Itoa(i)
			dirs[i] = fs.DirEntry(&DirEntry{
				info:   &FileInfo{name: base, isDir: isDir(v)},
				rootFS: fsys,
				path:   join(name, base)})
		}
	case map[string]interface{}:
		dirs = make([]fs.DirEntry, len(vv))
		names := make([]string, 0, len(vv))
		for n := range vv {
			names = append(names, n)
		}
		sort.Strings(names)
		for i, base := range names {
			dirs[i] = fs.DirEntry(&DirEntry{
				info:   &FileInfo{name: base, isDir: isDir(vv[base])},
				rootFS: fsys,
				path:   join(name, base)})
		}
	default:
		return nil, &fs.PathError{"readdir", name, fmt.Errorf("is not a directory")}
	}
	return dirs, nil
}
