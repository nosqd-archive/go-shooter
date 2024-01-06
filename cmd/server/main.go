package main

import (
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/nosqd/go-shooter/shooter/game"
	"github.com/nosqd/go-shooter/shooter/network"
	"github.com/nosqd/go-shooter/shooter/network/packets"
	"github.com/nosqd/go-shooter/shooter/network/packets/c2s"
	"io"
	"net"
)

var packetHandler = network.CreatePacketHandler()

func main() {
	config := game.ReadGameConfig("./client.ini")
	host := config.GetValue("network", "host").String()
	port, _ := config.GetValue("network", "port").Uint()

	server := network.GetNetwork()

	setup()
	server.ServerOnConnection(handleConnection)
	server.ServerListen(host, uint16(port))
}

func handleConnection(connection net.Conn) {
	cfmt.Printf("[%s] Network connected\n", connection.RemoteAddr().String())
	for {
		_, err := connection.Read(make([]byte, 0))
		if err != io.EOF && err != nil {
			break
		}

		packetHandler.Handle(connection)
	}
	cfmt.Printf("[%s] Network disconnected\n", connection.RemoteAddr().String())
}

func setup() {
	packetHandler.AddGlobalHandler(func(packet *network.Packet, connection net.Conn) {
		cfmt.Printf("[%s] Recieved a packet (size: %d, pid: %d)\n", connection.RemoteAddr().String(), packet.ContentLength, packet.PacketID)
	})

	packetHandler.AddHandler(int32(packets.C2SCLIENTHELLO), c2s.ClientHelloHandle)
}
