package utils

import (
	"encoding/json"
	"github/jaklove/framework/zinx/ziface"
	"io/ioutil"
)

/*
	存储一切有关Zinx框架的全局参数，供其他模块使用
	一些参数也可以通过 用户根据 zinx.json来配置
*/
type GlobalObj struct {
	TcpServer     ziface.IServer
	Host          string
	TcpPort       int    //当前服务器主机监听端口号
	Name          string //当前服务器名称
	Version       string //当前Zinx版本号
	MaxPacketSize uint32 //都需数据包的最大值
	MaxConn       int    //当前服务器主机允许的最大链接个数
}

/*
	定义一个全局的对象
*/
var GlobalObject *GlobalObj

//读取文件内容
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("/Users/zhourenjie/go-project/gp-zinx/conf/zinx.json")
	if err != nil {
		panic(err)
	}

	//将json数据解析到struct中
	//fmt.Printf("json :%s\n", data)
	err = json.Unmarshal(data, GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	//初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:          "ZinxServerApp",
		Version:       "V0.4",
		TcpPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxPacketSize: 4096,
	}

	//从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}
