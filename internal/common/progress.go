package common

import "fmt"


type Bar struct{
	percent int
	cur int64
	total int64
	rate string
	graph string
}

func (this *Bar)NewOption(start,total int64,tip string){
	this.cur=start
	this.total=total
	fmt.Println(tip)
	if this.graph==""{
		this.graph="█"
	}
}
//获取当前的百分比
func (this *Bar)getPercent()int{
	res:=float32(this.cur)/float32(this.total)*100
	return int(res)
}
//移动步长
func (this *Bar)Add(step int64){
	lastPosotion:=this.getPercent()
	this.cur+=step
	nowPosition:=this.getPercent()
	this.percent = nowPosition
	if lastPosotion!=nowPosition && nowPosition%2==0{
		this.rate+=this.graph
	}
	fmt.Printf("\r[%-50s]%3d%%  %8d/%d", this.rate, this.percent, this.cur, this.total)
}


func(this *Bar) Finish(){
	for i := 0; i < 50; i++ {
		this.rate+=this.graph
	}
	fmt.Printf("\r[%-50s]%3d%%  %8d/%d", this.rate, this.percent, this.cur, this.total)
	fmt.Println()
}
