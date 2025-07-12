package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/assets"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/weapons"
)

type Player struct {
	Position rl.Vector2
	Sprite   rl.Texture2D
	FrameRec rl.Rectangle
	Speed    float32
	Color    rl.Color
	width    float32
	height   float32
	Sword    weapons.Sword
}

func NewPlayer() Player {
	return Player{
		Position: rl.NewVector2(256, 112),
		Sprite:   assets.PlayerSprite,
		Speed:    200,
		Color:    rl.Blue,
		width:    16,
		height:   32,
		Sword:    weapons.NewSword(),
	}
}

func (p *Player) Update(camera rl.Camera2D) {
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
	p.Sword.UpdateSwordPos(p.Position, camera)
	p.Sword.SwordAttk(p.Sword)

}

func (p *Player) Draw() {
	rl.DrawTexture(p.Sprite, int32(p.Position.X), int32(p.Position.Y), rl.White)
	p.Sword.Draw()
}
