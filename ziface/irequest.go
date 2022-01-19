package ziface

/**
Irequest 接口
实际上是把客户端请求的链接信息 和请求的数据包装到了Request里面
*/
type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetData() []byte            //获取请求消息的数据
}
