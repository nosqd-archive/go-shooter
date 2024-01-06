package c2s

import (
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/nosqd/go-shooter/shooter/game"
	"github.com/nosqd/go-shooter/shooter/network"
	"net"
)

func ClientHelloCreate(name string) *network.Packet {
	packet := network.PacketCreate(int32(1001))

	packet.WriteUInt32(game.GAME_VERSION)
	packet.WriteString(name)

	return packet
}

var ClientHelloHandle = func(packet *network.Packet, connection net.Conn) {
	gameVersion := packet.ReadUInt32()
	playerName := packet.ReadString()

	cfmt.Printf("[%s] Network sent hello message\n", connection.RemoteAddr().String())
	cfmt.Printf("\tNetwork Version: %d\n", gameVersion)
	if gameVersion > game.GAME_VERSION {
		cfmt.Printf("\tServer is too old, disconnecting player\n")
		// TODO: send kicked message
		connection.Close()
	}
	if gameVersion > game.GAME_VERSION {
		cfmt.Printf("\tNetwork is too old, disconnecting player\n")
		// TODO: send kicked message
		connection.Close()
	}
	if gameVersion == game.GAME_VERSION {
		cfmt.Printf("\tNetwork Player name: %s\n", playerName)
		// TODO: sent serverhello packet
	}
}
