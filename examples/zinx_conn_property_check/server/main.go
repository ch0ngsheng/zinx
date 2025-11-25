package main

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

// OnConnStart 连接建立时的钩子函数
func OnConnStart(conn ziface.IConnection) {
	zlog.Ins().InfoF("[OnConnStart] Connection %s established", conn.RemoteAddrString())
}

// AuthRouter 处理认证消息的路由
type AuthRouter struct{}

// PreHandle 预处理
func (ar *AuthRouter) PreHandle(request ziface.IRequest) {
	zlog.Ins().InfoF("[AuthRouter] PreHandle: %s", string(request.GetData()))
}

// Handle 处理
func (ar *AuthRouter) Handle(request ziface.IRequest) {
	zlog.Ins().InfoF("[AuthRouter] Handle: %s", string(request.GetData()))

	// 解析认证消息
	name := string(request.GetData())

	// 检查name是否为test
	if name != "test" {
		zlog.Ins().ErrorF("[AuthRouter] Invalid name: %v, expected 'test'", name)
		// 断开连接
		request.GetConnection().Stop()
		return
	}

	zlog.Ins().InfoF("[AuthRouter] Connection %s passed authentication, name: %v", request.GetConnection().RemoteAddrString(), name)

	// 认证成功，发送认证通过消息
	err := request.GetConnection().SendMsg(2, []byte("auth_ok"))
	if err != nil {
		zlog.Ins().ErrorF("[AuthRouter] Failed to send auth_ok: %v", err)
	}
}

// PostHandle 后处理
func (ar *AuthRouter) PostHandle(request ziface.IRequest) {
	zlog.Ins().InfoF("[AuthRouter] PostHandle: %s", string(request.GetData()))
}

// OnConnStop 连接断开时的钩子函数
func OnConnStop(conn ziface.IConnection) {
	zlog.Ins().InfoF("[OnConnStop] Connection %s closed", conn.RemoteAddrString())
}

// PingRouter 处理ping消息的路由
type PingRouter struct{}

// PreHandle 预处理
func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	zlog.Ins().InfoF("[PingRouter] PreHandle: %s", string(request.GetData()))
}

// Handle 处理
func (pr *PingRouter) Handle(request ziface.IRequest) {
	zlog.Ins().InfoF("[PingRouter] Handle: %s", string(request.GetData()))
	// 回复pong
	err := request.GetConnection().SendMsg(1, []byte("pong"))
	if err != nil {
		zlog.Ins().ErrorF("[PingRouter] Failed to send pong: %v", err)
	}
}

// PostHandle 后处理
func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	zlog.Ins().InfoF("[PingRouter] PostHandle: %s", string(request.GetData()))
}

func main() {
	// 创建服务器实例
	server := znet.NewServer()

	// 设置连接钩子函数
	server.SetOnConnStart(OnConnStart)
	server.SetOnConnStop(OnConnStop)

	// 添加路由
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(255, &AuthRouter{})

	// 启动服务器
	fmt.Println("Starting server...")
	server.Serve()
}
