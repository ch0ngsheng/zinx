# Zinx Connection Property Check Demo

This demo shows how to validate connection properties in Zinx framework.

## Description

The server will validate the connection property `name` when a client connects. If the `name` is not `test`, the server will close the connection.

## How to run

### 1. Start the server

```bash
cd server
 go run main.go
```

### 2. Start the client

```bash
cd client
 go run main.go
```

## Expected output

### Server

```
[START] TCP Server name: ZinxServer,listener at IP: 0.0.0.0, Port 8999 is starting
[OnConnStart] Connection 127.0.0.1:12345 established
[OnConnStart] Connection 127.0.0.1:12345 passed validation, name: test
[PingRouter] PreHandle: ping
[PingRouter] Handle: ping
[PingRouter] PostHandle: ping
```

### Client

```
Starting client...
[START] Zinx Client dial RemoteAddr: 127.0.0.1:8999
[START] Zinx Client LocalAddr: 127.0.0.1:12345, RemoteAddr: 127.0.0.1:8999
[OnConnStart] Connection to 127.0.0.1:8999 established
[PongRouter] PreHandle: pong
[PongRouter] Handle: pong
[PongRouter] PostHandle: pong
Client exiting after 5 seconds
[STOP] Zinx Client LocalAddr: 127.0.0.1:12345, RemoteAddr: 127.0.0.1:8999
[OnConnStop] Connection to 127.0.0.1:8999 closed
```

## Test invalid name

To test invalid name, modify the client code in `client/main.go`:

```go
// 设置连接属性name=invalid
conn.SetProperty("name", "invalid")
```

Then run the client again. The server will close the connection immediately.

### Server output for invalid name

```
[OnConnStart] Connection 127.0.0.1:12346 established
[OnConnStart] Invalid name property: invalid, expected 'test'
[OnConnStop] Connection 127.0.0.1:12346 closed
```

### Client output for invalid name

```
Starting client...
[START] Zinx Client dial RemoteAddr: 127.0.0.1:8999
[START] Zinx Client LocalAddr: 127.0.0.1:12346, RemoteAddr: 127.0.0.1:8999
[OnConnStart] Connection to 127.0.0.1:8999 established
[OnConnStop] Connection to 127.0.0.1:8999 closed
```
