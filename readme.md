## 命令行-mytool

## 1. 目的

将项目参数化,形成一个本地的脚手架库

## 2. 原理

将现有的项目通过`挖空`的方式,参数化,比如xxx公司,可以使用`{{companyName}}`来代替,并将这个配置到最外层的settings.json中,通过命令执行字段的替换

## 3. 使用

如果想直接使用的话,可以在`release`文件夹下对应的包,让后将`mytool`添加到环境变量即

下面是几个已经做好的命令

```shell
mytool new
```

生成一个固定配置文件和目录结构

```shell
|--template 模板文件的存放地点
|--output   将字段替换好的文件
|--settings.json 配置文件
```

下面是settings.json的配置文件说明

```go
{
  "template": "", //模板文件目录
  "output": "", //文件输出地址
  "worker_num": 10,//执行时协程数量
    "separator": "",//分隔符比如{{}}
  "fields": {// 用于替换的字段
  }
}
```

后期会增加`mytool list`和`mytool install` ,install将已经做好的模板库,安装到中心库,以后可以直接使用list查看,并直接使用,不必再次替换