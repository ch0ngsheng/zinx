package main

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

// DoConnectionBegin 连接建立时执行的函数
func DoConnectionBegin(conn ziface.IConnection) {
	zlog.Ins().InfoF("DoConnectionBegin is Called ...")
}

// DoConnectionLost 连接断开时执行的函数
func DoConnectionLost(conn ziface.IConnection) {
	zlog.Ins().InfoF("Conn is Lost")
}

func main() {
	// 创建服务器
	s := znet.NewServer()

	// 设置连接钩子函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	// 添加路由处理认证消息
	s.AddRouter(1, &AuthRouter{})

	// 启动服务器
	s.Serve()
}

// AuthRouter 处理认证消息的路由
// (处理认证消息的路由)
type AuthRouter struct {
	znet.BaseRouter
}

// Handle handles the authentication message
// (处理认证消息)
func (r *AuthRouter) Handle(req ziface.IRequest) {
	// Get the connection object
	conn := req.GetConnection()

	// Get the message data
	data := req.GetData()

	// Parse the message data to get the name property
	name := string(data)

	// Check if the name property is valid
	if name != "test" {
		zlog.Ins().InfoF("Connection auth failed: invalid name=%s. Closing connection...", name)
		conn.Stop()
		return
	}

	zlog.Ins().InfoF("Connection auth success: name=%s", name)

	// Send an authentication success message to the client
	if err := conn.SendMsg(2, []byte("Auth Success")); err != nil {
		zlog.Ins().InfoF("Failed to send auth success message: %v", err)
	}
}
