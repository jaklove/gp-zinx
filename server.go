package main

import (
	"fmt"
	"github/jaklove/framework/zinx/ziface"
	"github/jaklove/framework/zinx/znet"
)

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

////Test PreHandle
//func (this *PingRouter) PreHandle(request ziface.IRequest) {
//	fmt.Println("Call Router PreHandle")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
//	if err != nil {
//		fmt.Println("call back ping ping ping error")
//	}
//}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//回写数据
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

//func (this *PingRouter) PostHandle(request ziface.IRequest) {
//	fmt.Println("Call Router PostHandle")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
//	if err != nil {
//		fmt.Println("call back ping ping ping error")
//	}
//}

func main() {
	//1 创建一个server 句柄 s
	s := znet.NewServer("[zinx V0.1]")

	s.AddRouter(&PingRouter{})
	// 开启服务
	s.Serve()
}
