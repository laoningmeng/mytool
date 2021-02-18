package code

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
)
func CheckExist(path string)bool{
	_, err := os.Stat(path)
	if err!=nil{
		return false
	}
	return true
}

func CreateFolder(path string)error{
	if exist:=CheckExist(path);!exist{
		//不存在创建文件夹
		e := os.Mkdir(path, os.ModePerm)
		return e
	}
	return nil
}

//创建文件
func CreateFile(path string,content []byte) error{
	file,err:=os.Create(path)
	defer file.Close()
	cobra.CheckErr(err)
	_,e:=file.Write(content)
	cobra.CheckErr(e)
	return nil
}
//在当前文件夹创建文件
func CreateCurrentFile(name string,Content []byte)error{
	path,e:=os.Getwd()
	cobra.CheckErr(e)
	name = path+"/"+name
	return CreateFile(name,Content)
}
func CreateCurrentFolder(path string)error{
	current,e:=os.Getwd()
	cobra.CheckErr(e)
	path=current+"/"+path
	return CreateFolder(path)
}

//在home下创建文件
func CreateHomeFile(name string,Content []byte)error{
	path, err := homedir.Dir()
	cobra.CheckErr(err)
	name = path+"/"+name
	return CreateFile(name,Content)
}

func CreateHomeFolder(path string)error{
	current, err := homedir.Dir()
	cobra.CheckErr(err)
	path=current+"/"+path
	return CreateFolder(path)
}

