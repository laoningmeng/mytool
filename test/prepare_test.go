package test

import (
	"mytool/internal/code"
	"testing"
)

func TestCheckExist(t *testing.T) {

}

func TestCreateFolder(t *testing.T) {
	path:="test"
	code.CreateFolder(path)
}

func TestCreateCurrentFolder(t *testing.T) {
	path:="ojbk"
	code.CreateCurrentFolder(path)
}

func TestCreateCurrentFile(t *testing.T) {
	code.CreateCurrentFile("ojbk.txt",[]byte("hello,owrld"))
}