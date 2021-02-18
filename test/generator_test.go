package test

import (
	"fmt"
	"mytool/internal/code"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	start:=time.Now()
	currentDir:=`C:\workspace\git_repo\vue\vue-electron-template`
	targetDir := `C:\Users\lx\Desktop\copytest\output`
	g:=code.Genrator{
		Src:currentDir,
		Dst:targetDir,
	}
	g.Run()
	fmt.Println("耗时",time.Since(start))
}

func TestReadSettings(t *testing.T) {
	fields:=code.ReadSettings()
	fmt.Println(fields)
}