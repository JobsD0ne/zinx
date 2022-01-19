package znet

import (
	"errors"
	"fmt"
	"net"
	"time"
	"zinx/ziface"
)

type Server struct {
	//服务器名称
	Name string
	//tcp4 or other
	IpVersion string
	//服务绑定的ip地址
	IP string
	//端口号
	Port int
	//当前Server由用户 绑定的回调router，也就是Server注册的链接对应的处理业务
	Router ziface.IRouter
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[conn handle ] CallbackToClient ...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("callBackToClient error")
	}
	return nil
}
func (s *Server) Start() {
	fmt.Printf("[srart]Server listenner at IP: %s, Port :%d ,is starting \n", s.IP, s.Port)
	go func() {
		//1 获取一个Tcp的addr
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}
		//2 监听服务器地址
		listenner, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IpVersion, "err ", err)
			return
		}
		fmt.Println("start Zinx Server", s.Name, "successfully ,now linstenning....")
		//TODO server.go 应该有一个自动生成ID的方法
		var cid uint32
		cid = 0

		//启动server网络连接服务
		for {
			//3.1阻塞等待客户端建立连接请求
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accecp err", err)
				continue
			}
			//3.2TODO server.Start()设置服务器最大连接控制，如果超过最大连接，那么则关闭此薪的连接

			//3.3 处理该新连接的业务方法，此时应该有handle和conn是绑定的
			dealConn := NewConntion(conn, cid, CallBackToClient)
			cid++
			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()
}
func (s *Server) Stop() {
	fmt.Println("[STOP] Zin server ,name", s.Name)

}
func (s *Server) Serve() {
	s.Start()
	//阻塞，否则主go推出，listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}

}

//路由功能，给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("add router successfully!")
}
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IpVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
		Router:    nil,
	}
	return s
}
