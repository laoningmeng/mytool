package test

import (
	"fmt"
	"mytool/internal/code"
	"testing"
	"time"
)

func TestCopyDir(t *testing.T) {

	currentDir:=`C:\workspace\git_repo\vue\vue-electron-template`
	targetDir := `C:\Users\lx\Desktop\copytest\output`
	//_=
	outter:=time.Now()
	fmt.Println(currentDir,targetDir)
	code.Dir(currentDir, targetDir)
	outterEnd:=time.Since(outter)
	fmt.Println("总共耗时:",outterEnd)
}