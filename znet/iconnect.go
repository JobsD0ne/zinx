package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	//当前链接的socket tcp套接字
	Conn *net.TCPConn
	//当前链接的ID，也可以称作session ID，id全局唯一
	ConnID uint32
	//当前链接的关闭状态
	isClosed bool
	//该链接的处理方法api
	handAPI ziface.HandFunc
	//告知该链接已经退出/停止的channle
	ExitBuffChan chan bool
	//该连接处理方法router
	Router ziface.IRouter
}

//创建连接的方法
func NewConntion(conn *net.TCPConn, connID uint32, callback_api ziface.HandFunc) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		handAPI:      callback_api,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running ")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()
	for {
		//读取我们最大的数据到buf中
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- true
			continue
		}
		fmt.Println(buf)
		//得到当前客户端请求的Request数据
		req := Request{
			conn: c,
			data: buf,
		}
		//从路由Routers 中找到注册绑定Conn的对应Handle
		go func(request ziface.IRequest) {
			//执行注册的路由方法
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

//启动连接，让连接开始工作
func (c *Connection) Start() {
	//开启处理该链接读取到客户端数据之后的请求业务
	go c.StartReader()
	// for {
	// 	select {

	// 	case <-c.ExistBuffChan:
	// 		//得到退出消息不再阻塞
	// 		return
	// 	}
	// }
	for range c.ExitBuffChan {
		return
	}
}
func (c *Connection) Stop() {
	if c.isClosed {
		return
	}

	c.isClosed = true
	//TODO Connection STOP() 如果用户注册了该链接的关闭回调服务，那么在此刻应该显示调用

	//关闭socket连接
	c.Conn.Close()
	//通知从缓存队列度数据的业务，该链接已经关闭
	c.ExitBuffChan <- true
	//关闭该链接全部管道
	close(c.ExitBuffChan)
}

//从当前的连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端地址信息
func (c *Connection) GetAddrInfo() net.Addr {
	return c.Conn.RemoteAddr()
}
