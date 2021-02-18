package code

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"mytool/internal/common"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type SourceQueue struct{
	Src string
	Dst string
	IsDir bool
}

type ConcurrencyGenerator struct{
	WorkerNum int
	Src string
	Dst string
	FilesQueue []SourceQueue
	FoldersQueue []SourceQueue
	Settings map[string]interface{}
}

func (this *ConcurrencyGenerator) Run(){
	//读取配置文件,将正则的规则载入
	this.GetSettings()
	//拿到所有的path
	//首先将读取的目录结构放进队列
	this.IterateFolders(this.Src,this.Dst)
	length:=len(this.FilesQueue)
	in:=make(chan SourceQueue)
	processNum:=make(chan int) //处理文件的
	isFinished:=make(chan bool) //是否处理完成
	//设置计数器
	counter(processNum,length,isFinished)

	for i := 0; i < this.WorkerNum; i++ {
		this.createWorker(in,processNum)//设置接收者
	}

	//发送消息
	var bar common.Bar
	bar.NewOption(0,int64(length),"拷贝文件")
	for _,v:=range this.FilesQueue{
		in<-v
		bar.Add(1)
	}
	<-isFinished
	bar.Finish()
}
func counter(c chan int, total int,finish chan bool){
	count:=0
	go func() {
		for{
			<-c
			count++
			if count>=total{
				finish<-true
				break
			}
		}
	}()
}

func (this *ConcurrencyGenerator) createWorker(in chan SourceQueue,c chan int){
	go func() {
		for{
			data:=<-in
			err:=this.worker(data)
			if err!=nil{
				fmt.Println("Error:",err)
				continue
			}
			c<-1
		}
	}()
}

func (this *ConcurrencyGenerator)worker(data SourceQueue)(error){
	if data.IsDir{
		err:=os.MkdirAll(data.Dst, os.ModePerm)
		return err
	}else{
		f, err := os.OpenFile(data.Src, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Println("[err1]:",err)
			return err
		}
		bytes, err := ioutil.ReadAll(f)
		if err!=nil{
			fmt.Println("[err2]:",err)
			return err
		}
		content:=string(bytes)
		fields:=this.Settings["fields"]
		sep:=this.Settings["separator"].(string)
		var res string
		for index,value:=range fields.(map[string]interface{}){
			word:=sep+index+sep
			res=strings.ReplaceAll(content,word,value.(string))
		}
		//写入之前要进行文件夹是否已经创建的检测:
		exist, _ := PathExists(data.Dst)
		if !exist{
			dir, _ := filepath.Split(data.Dst)
			_=os.MkdirAll(dir, os.ModePerm)
		}
		err = ioutil.WriteFile(data.Dst, []byte(res), os.ModePerm)
		if err!=nil{
			fmt.Println("[err4]:",err)
			return err
		}
	}
	return nil
}


// Dir copies a whole directory recursively
func (this *ConcurrencyGenerator)IterateFolders(src string, dst string) {
	var err error
	var fds []os.FileInfo
	if fds, err = ioutil.ReadDir(src); err != nil {
		cobra.CheckErr(err)
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())
		if fd.IsDir() {
			this.FilesQueue=append(this.FilesQueue,SourceQueue{
				Src:srcfp,
				Dst:dstfp,
				IsDir: true,
			})
			this.IterateFolders(srcfp, dstfp)
		} else {
			this.FilesQueue=append(this.FilesQueue,SourceQueue{
				Src:srcfp,
				Dst:dstfp,
				IsDir: false,
			})
		}
	}
}

func (this *ConcurrencyGenerator)GetSettings(){
	position:=common.GetCurrentPosition()
	filePath:=position
	viper.SetConfigName("settings")
	viper.SetConfigType("json")
	viper.AddConfigPath(filePath)
	_=viper.ReadInConfig()
	settings:=viper.AllSettings()
	//检测两个必须存在的字段,没有就结束,不再向后进行
	if _,ok:=settings["fields"];!ok{
		cobra.CheckErr(errors.New("settings.json中缺少fields字段"))
	}
	if _,ok:=settings["separator"];!ok{
		cobra.CheckErr(errors.New("settings.json中缺少separator字段"))
	}
	if _,ok:=settings["template"];!ok{
		cobra.CheckErr(errors.New("settings.json中缺少template字段"))
	}
	if _,ok:=settings["output"];!ok{
		cobra.CheckErr(errors.New("settings.json中缺少output字段"))
	}
	this.Src=viper.GetString("template")
	this.Dst=viper.GetString("output")
	this.WorkerNum=viper.GetInt("worker_num")
	this.Settings=settings
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
