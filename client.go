package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPost int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPost,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPost))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}
	client.conn = conn
	return client
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊")
	fmt.Println("2.私聊")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")

	// 命令行解析
	flag.Parse()
}

func main() {
	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println("》》》》》》》服务器连接失败")
		return
	}
	fmt.Println(">>>>>>服务器连接成功")
	select {}
}
