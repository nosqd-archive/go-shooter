package game

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	"github.com/nosqd/go-shooter/shooter/network"
)

var gameInstance *Game = nil

type Game struct {
	client *network.Network
}

func GetGame() *Game {
	if gameInstance == nil {
		game := Game{
			client: network.GetNetwork(),
		}
		gameInstance = &game
	}

	return gameInstance
}

func (*Game) Update() {

}

func (self *Game) Draw() {

	DrawText(
		fmt.Sprintf("Shooter Game v%d (%d FPS)", GAME_VERSION, GetFPS()),
		5,
		5,
		18,
		RayWhite,
	)
}
