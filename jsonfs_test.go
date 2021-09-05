package jsonfs

import (
	"fmt"
	"testing"
)

const (
	dataStr = `{"menu": {
  "id": "file",
  "value": "File",
  "popup": {
    "menuitem": [
      {"value": "New", "onclick": "CreateNewDoc()"},
      {"value": "Open", "onclick": "OpenDoc()"},
      {"value": "Close", "onclick": "CloseDoc()"}
    ]
  }
}}`
)

func TestNamev(t *testing.T) {
	fsys, _ := NewFS([]byte(dataStr))
	fmt.Println(fsys.namev(""))
	fmt.Println(fsys.namev("."))
	fmt.Println(fsys.namev("menu/id/ok"))
	fmt.Println(fsys.namev("menu/popup/menuitem/2/onclick"))
	fmt.Println(fsys.namev("menu/popup/menuitem/3/value"))
}

func TestReadDir(t *testing.T) {
	fsys, _ := NewFS([]byte(dataStr))
	fmt.Println(fsys.ReadDir(""))
	fmt.Println(fsys.ReadDir("."))
	fmt.Println(fsys.ReadDir("menu"))
	fmt.Println(fsys.ReadDir("menu/popup/menuitem/2/onclick"))
	fmt.Println(fsys.ReadDir("menu/popup/menuitem/3"))
}
