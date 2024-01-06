package network

import (
	"github.com/i582/cfmt/cmd/cfmt"

	"io"
	"net"
	"os"
)

var networkInstance *Network

type Network struct {
	Host string
	Port uint16

	connection           *net.Conn
	packetHandler        *PacketHandler
	onSetupCallback      func()
	onConnectionCallback func(conn net.Conn)
}

func GetNetwork() *Network {
	if networkInstance == nil {
		client := Network{
			packetHandler: CreatePacketHandler(),
		}
		networkInstance = &client
	}

	return networkInstance
}

func (self *Network) GetPacketHandler() *PacketHandler {
	return self.packetHandler
}

func (self *Network) ClientConnect(Host string, Port uint16) {
	self.Host = Host
	self.Port = Port
	connection, err := net.Dial("tcp", cfmt.Sprintf("%s:%d", self.Host, self.Port))
	self.connection = &connection
	if err != nil {
		cfmt.Printf("{{Fail to connect: %v.}}::red|bold\n", err)
		os.Exit(1)
	}
	cfmt.Printf("{{Connected to }}::green{{%s:%d}}::green|bold\n", self.Host, self.Port)

}

func (self *Network) ClientConnection() *net.Conn {
	return self.connection

}

func (self *Network) ClientWorkRoutine() {
	_, err := (*self.ClientConnection()).Read(make([]byte, 0))
	if err != io.EOF && err != nil {
		return
	}

	self.packetHandler.Handle(*self.ClientConnection())
}

func (self *Network) ClientClose() {
	(*self.ClientConnection()).Close()
	cfmt.Printf("{{Goodbye.}}::green|bold\n")
}

func (self *Network) ServerOnConnection(handler func(connection net.Conn)) {
	self.onConnectionCallback = handler
}

func (self *Network) ServerListen(Host string, Port uint16) {
	self.Host = Host
	self.Port = Port
	cfmt.Printf("Starting on {{%s:%d}}::green|bold\n", self.Host, self.Port)

	ln, err := net.Listen("tcp", cfmt.Sprintf("%s:%d", self.Host, self.Port))

	if err != nil {
		cfmt.Printf("{{Fail to start listening: %v.}}::red|bold\n", err)
		os.Exit(1)
	}
	cfmt.Printf("{{Working on }}::green{{%s:%d}}::green|bold\n", self.Host, self.Port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			cfmt.Printf("{{Fail to accept connection, continuing work: %v.}}::yellow|bold\n", err)
		}
		go self.onConnectionCallback(conn)
	}
}
