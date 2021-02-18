package code

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"
)

// File copies a single file from src to dst
func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		cobra.CheckErr(err)
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		cobra.CheckErr(err)
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		cobra.CheckErr(err)
	}
	if srcinfo, err = os.Stat(src); err != nil {
		cobra.CheckErr(err)
	}
	return os.Chmod(dst, srcinfo.Mode())
}


// Dir copies a whole directory recursively
func Dir(src string, dst string) error {
	start := time.Now() // 获取当前时间
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		cobra.CheckErr(err)
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		cobra.CheckErr(err)
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		cobra.CheckErr(err)
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			Dir(srcfp, dstfp)
		} else {
			//File(srcfp, dstfp)
			fmt.Println(srcfp)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
	return nil
}