package network

import (
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/nosqd/go-shooter/shooter/packets"
	c2s2 "github.com/nosqd/go-shooter/shooter/packets/c2s"
	"github.com/nosqd/go-shooter/shooter/packets/s2c"
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

func (self *PacketHandler) _setupGlobal() {
	self.AddGlobalHandler(func(packet *Packet, connection net.Conn) {
		cfmt.Printf("[%s] Recieved a packet (size: %d, pid: %d)\n", connection.RemoteAddr().String(), packet.ContentLength, packet.PacketID)
	})
}

func (self *PacketHandler) SetupClient() {
	self._setupGlobal()
	self.AddHandler(int32(packets.S2CSERVERHELLO), s2c.ServerHelloHandle)
}

func (self *PacketHandler) SetupServer() {
	self._setupGlobal()
	self.AddHandler(int32(packets.C2SCLIENTHELLO), c2s2.ClientHelloHandle)
	self.AddHandler(int32(packets.C2SPLAYERUDPATE), c2s2.PlayerUpdateHandle)
}
