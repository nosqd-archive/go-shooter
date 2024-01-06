package shooter

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"github.com/nosqd/go-shooter/shooter/game"
	"github.com/nosqd/go-shooter/shooter/packets/c2s"
	"github.com/nosqd/go-shooter/utils"
)

type Player struct {
	PlayerID uint32
	Position Vector2
	Name     string
	IsLocal  bool

	lastSentPosition Vector2
}

func PlayerCreate(id uint32, pos Vector2, name string, isLocal bool) *Player {
	p := Player{}
	p.PlayerID = id
	p.Position = pos
	p.Name = name
	p.IsLocal = isLocal
	p.lastSentPosition = pos

	return &p
}

func (self *Player) Update(game *game.Game) {
	self.Position.X = GetMousePosition().X
	self.Position.Y = GetMousePosition().Y

	if Vector2Distance(self.Position, self.lastSentPosition) > 20 {
		self.lastSentPosition = self.Position
		packet := c2s.PlayerUpdateCreate(self.PlayerID, self.Position)
		packet.Write(*game.Client.ClientConnection())
	}
}

func (self *Player) Draw(game *game.Game) {
	DrawCircle(
		int32(self.Position.X),
		int32(self.Position.Y),
		10,
		utils.Uint32ToColor(self.PlayerID),
	)
}
