package network

import (
	"net"
)

type PacketHandlerCallback func(packet *Packet, connection net.Conn)

type PacketHandler struct {
	handlers map[int32][]PacketHandlerCallback
}

func CreatePacketHandler() *PacketHandler {
	handler := PacketHandler{}

	handler.handlers = map[int32][]PacketHandlerCallback{}

	return &handler
}

func (self *PacketHandler) AddHandler(packetId int32, callback PacketHandlerCallback) {
	self.handlers[packetId] = append(self.handlers[packetId], callback)
}

func (self *PacketHandler) AddGlobalHandler(callback PacketHandlerCallback) {
	self.handlers[-1] = append(self.handlers[-1], callback)
}

func (self *PacketHandler) Handle(connection net.Conn) {
	packet := ReadPacket(connection)

	if packet == nil {
		connection.Close()
		return
	}

	for _, g := range self.handlers[-1] {
		g(packet, connection)
	}

	for _, g := range self.handlers[packet.PacketID] {
		g(packet, connection)
	}
}
