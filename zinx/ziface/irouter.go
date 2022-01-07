package ziface

/*
	路由接口， 这里面路由是 使用框架者给该链接自定的 处理业务方法
	路由里的IRequest 则包含用该链接的链接信息和该链接的请求数据信息
*/
type IRouter interface {
	PreHandle(request IRequest)   //在处理conn业务之前的钩子方法
	Handle(request IRequest)	 //处理conn业务的方法
	PostHandle(request IRequest) //处理conn业务之后的钩子方法
}

type Iserver interface {
	//启动服务器方法
	Start()
	//停止服务器方法
	Stop()
	//开启业务服务方法
	Serve()
	//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(router IRouter)
}

