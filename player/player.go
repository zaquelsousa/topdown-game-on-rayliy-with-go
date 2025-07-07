package player

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position rl.Vector2
	Sprite   rl.Texture2D
	FrameRec rl.Rectangle
	Speed    float32
	Color    rl.Color
	width    float32
	height   float32
}

const (
	FrameCount      = 2
	FrameWidthRatio = 1.0 / FrameCount
)

func NewPlayer(sprite rl.Texture2D) Player {
	return Player{
		Position: rl.NewVector2(256, 112),
		Sprite:   sprite,
		Speed:    200,
		Color:    rl.Blue,
		width:    16,
		height:   32,
	}
}

func (p *Player) Update() {
	if rl.IsKeyDown(rl.KeyD) {
		p.Position.X += p.Speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.Position.X -= p.Speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyW) {
		p.Position.Y -= p.Speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Position.Y += p.Speed * rl.GetFrameTime()
	}
}

func (p *Player) Draw() {
	rl.DrawTexture(p.Sprite, int32(p.Position.X), int32(p.Position.Y), rl.White)
}
