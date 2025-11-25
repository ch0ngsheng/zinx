package main

import (
	"fmt"
	"time"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

// OnConnStart 连接建立时的钩子函数
func OnConnStart(conn ziface.IConnection) {
	zlog.Ins().InfoF("[OnConnStart] Connection to %s established", conn.RemoteAddrString())

	// 发送认证消息
	err := conn.SendMsg(255, []byte("test"))
	if err != nil {
		zlog.Ins().ErrorF("[OnConnStart] Failed to send auth message: %v", err)
	}
}

// AuthOkRouter 处理认证通过消息的路由
type AuthOkRouter struct{}

// PreHandle 预处理
func (aor *AuthOkRouter) PreHandle(request ziface.IRequest) {
	zlog.Ins().InfoF("[AuthOkRouter] PreHandle: %s", string(request.GetData()))
}

// Handle 处理
func (aor *AuthOkRouter) Handle(request ziface.IRequest) {
	zlog.Ins().InfoF("[AuthOkRouter] Handle: %s", string(request.GetData()))

	// 认证通过，发送ping消息
	err := request.GetConnection().SendMsg(0, []byte("ping"))
	if err != nil {
		zlog.Ins().ErrorF("[AuthOkRouter] Failed to send ping: %v", err)
	}
}

// PostHandle 后处理
func (aor *AuthOkRouter) PostHandle(request ziface.IRequest) {
	zlog.Ins().InfoF("[AuthOkRouter] PostHandle: %s", string(request.GetData()))
}

// OnConnStop 连接断开时的钩子函数
func OnConnStop(conn ziface.IConnection) {
	zlog.Ins().InfoF("[OnConnStop] Connection to %s closed", conn.RemoteAddrString())
}

// PongRouter 处理pong消息的路由
type PongRouter struct{}

// PreHandle 预处理
func (pr *PongRouter) PreHandle(request ziface.IRequest) {
	zlog.Ins().InfoF("[PongRouter] PreHandle: %s", string(request.GetData()))
}

// Handle 处理
func (pr *PongRouter) Handle(request ziface.IRequest) {
	zlog.Ins().InfoF("[PongRouter] Handle: %s", string(request.GetData()))
}

// PostHandle 后处理
func (pr *PongRouter) PostHandle(request ziface.IRequest) {
	zlog.Ins().InfoF("[PongRouter] PostHandle: %s", string(request.GetData()))
}

func main() {
	// 创建客户端实例
	client := znet.NewClient("127.0.0.1", 8999)

	// 设置连接钩子函数
	client.SetOnConnStart(OnConnStart)
	client.SetOnConnStop(OnConnStop)

	// 添加路由
	client.AddRouter(1, &PongRouter{})
	client.AddRouter(2, &AuthOkRouter{})

	// 启动客户端
	fmt.Println("Starting client...")
	client.Start()

	// 保持客户端运行
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Client exiting after 5 seconds")
		client.Stop()
	}
}
