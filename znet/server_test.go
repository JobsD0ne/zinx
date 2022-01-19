package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("Client test ....start")
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client Dial err, ", err)
		return
	}
	for {
		_, err := conn.Write([]byte("hello ZINX"))
		if err != nil {
			fmt.Println("write error err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error", err)
			return
		}
		fmt.Printf(" Server call back :%s,cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}
func TestServer(t *testing.T) {
	//服务端测试
	s := NewServer("[zinx V0.1]")

	//客户端测试
	go ClientTest()

	//2开启服务
	s.Serve()
}
