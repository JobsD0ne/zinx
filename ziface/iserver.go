package ziface

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//开启业务服务方法
	Serve()
	//路由功能，给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(router IRouter)
}
