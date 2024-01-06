package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/nosqd/go-shooter/shooter/game"
	"github.com/nosqd/go-shooter/shooter/network"
	"github.com/nosqd/go-shooter/shooter/network/packets/c2s"
	"net"
)

func main() {
	config := game.ReadGameConfig("./client.ini")
	host := config.GetValue("network", "host").String()
	port, _ := config.GetValue("network", "port").Uint()
	playerName := config.GetValue("game", "player-name").String()

	cfmt.Printf("Starting on {{%s:%d}}::green|bold\n", host, port)

	client := network.GetNetwork()
	game := game.GetGame()

	client.ClientConnect(host, uint16(port))

	helloPacket := c2s.ClientHelloCreate(playerName)
	helloPacket.Write(*client.ClientConnection())

	setup(client)

	go client.ClientWorkRoutine()
	InitWindow(1280, 720, "game")

	for !WindowShouldClose() {
		game.Update()

		BeginDrawing()
		ClearBackground(Black)

		game.Draw()

		EndDrawing()
	}

	client.ClientClose()
	cfmt.Println("{{Goodbye}}::green|bold")
}

func setup(client *network.Network) {
	client.GetPacketHandler().AddGlobalHandler(func(packet *network.Packet, connection net.Conn) {
		cfmt.Printf("[%s] Recieved a packet (size: %d, pid: %d)\n", connection.RemoteAddr().String(), packet.ContentLength, packet.PacketID)
	})
}
