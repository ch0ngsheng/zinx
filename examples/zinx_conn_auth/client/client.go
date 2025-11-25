package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

// DoClientConnectedBegin 连接建立时执行的函数
func DoClientConnectedBegin(conn ziface.IConnection) {
	zlog.Debug("DoClientConnectedBegin is Called ... ")

	fmt.Println("Client connected to server")
	// Send an authentication message to the server
	if err := conn.SendMsg(1, []byte("test")); err != nil {
		zlog.Debug("Failed to send auth message: %v", err)
	}
}

// DoClientConnectedLost 连接断开时执行的函数
func DoClientConnectedLost(conn ziface.IConnection) {
	fmt.Println("Client disconnected from server")
}

// customRouter handles messages from the server
// (处理服务器发送的消息)
type customRouter struct {
	znet.BaseRouter
}

// Handle handles the message
// (处理消息)
func (r *customRouter) Handle(req ziface.IRequest) {
	// Get the message data
	data := req.GetData()

	// Print the message data
	zlog.Ins().InfoF("Received message from server: %s", string(data))
}

func main() {
	// 创建客户端
	client := znet.NewClient("127.0.0.1", 8999)

	// 设置连接建立时的钩子函数
	client.SetOnConnStart(DoClientConnectedBegin)

	// 设置连接断开时的钩子函数
	client.SetOnConnStop(DoClientConnectedLost)

	// 添加路由处理函数
	client.AddRouter(1, &customRouter{})

	// 添加处理认证成功消息的路由
	client.AddRouter(2, &customRouter{})

	// 启动客户端
	client.Start()

	// 等待中断信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	// 停止客户端
	client.Stop()
}
