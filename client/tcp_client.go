package client

import (
	"go-chat-server-practice/protocol"
	"net"
)

// 实现 ChatClient 接口
type TcpChatClient struct {
	name      string   // 客户端名称
	conn      net.Conn // 与服务端的连接
	cmdReader *protocol.CommandReader
	cmdWriter *protocol.CommandWriter
	incoming  chan protocol.MessageCommand // 服务端广播的消息会发送到这个 channel
}

func NewClient() *TcpChatClient {
	return &TcpChatClient{
		incoming: make(chan protocol.MessageCommand),
	}
}

// 连接到服务端。这里需要做一些初始化的工作
func (c *TcpChatClient) Dial(address string) error {
	// TODO
	return nil
}

// 获取服务端发来的命令，执行相应的操作
func (c *TcpChatClient) Start() {
	// TODO
}

// 简单地关闭连接
func (c *TcpChatClient) Close() {
	_ = c.conn.Close()
}

// 向服务端发送一条命令
func (c *TcpChatClient) Send(command interface{}) error {
	return c.cmdWriter.Write(command)
}

func (c *TcpChatClient) SetName(name string) error {
	return c.Send(protocol.NameCommand{name})
}

func (c *TcpChatClient) SendMessage(message string) error {
	return c.Send(protocol.SendCommand{message})
}

func (c *TcpChatClient) Incoming() chan protocol.MessageCommand {
	return c.incoming
}
