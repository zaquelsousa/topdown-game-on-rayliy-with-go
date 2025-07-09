package weapons

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/assets"
)

type Sword struct {
	Pos          rl.Vector2
	Width        float32
	Height       float32
	Damege       int32
	IsAttk       bool
	AttkTimer    float32
	AttkDuration float32
	rotation     float32
	startAngle   float32
	endAngle     float32
	baseAngle    float32
	Texture      rl.Texture2D
}

var SwordSprite = rl.Texture2D{}

func NewSword() Sword {
	return Sword{
		Pos:          rl.NewVector2(256, 112),
		Width:        4,
		Height:       12,
		Damege:       15,
		IsAttk:       false,
		AttkTimer:    0,
		AttkDuration: 1,
		Texture:      assets.SwordSprite,
	}
}

func (s *Sword) SwordAttk() {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && !s.IsAttk {
		s.IsAttk = true
		s.AttkTimer = 0

		s.baseAngle = s.rotation
		s.startAngle = s.baseAngle - 45
		s.endAngle = s.baseAngle + 45
	}
}

func (s *Sword) UpdateSwordPos(playerPos rl.Vector2, camera rl.Camera2D) {
	mouseScreenPos := rl.GetMousePosition()
	mouseWorldPos := rl.GetScreenToWorld2D(mouseScreenPos, camera)

	playerCenter := rl.NewVector2(playerPos.X+8, playerPos.Y+16)

	direction := rl.Vector2Subtract(mouseWorldPos, playerCenter)
	direction = rl.Vector2Normalize(direction)

	offset := float32(16)

	if !s.IsAttk {
		s.Pos = rl.Vector2Add(playerCenter, rl.Vector2Scale(direction, offset))
		angle := rl.Vector2Angle(rl.NewVector2(0, -1), direction) * (180.0 / 3.14159)
		s.rotation = angle
	} else {
		s.AttkTimer += rl.GetFrameTime()
		t := s.AttkTimer / s.AttkDuration
		if t > 1 {
			t = 1
			s.IsAttk = false
		}

		s.rotation = s.startAngle + (s.endAngle-s.startAngle)*t

		rad := float64(s.rotation * (math.Pi / 180.0))
		dir := rl.NewVector2(float32(math.Sin(rad)), -float32(math.Cos(rad)))
		s.Pos = rl.Vector2Add(playerCenter, rl.Vector2Scale(dir, offset))
	}
}

func (s *Sword) Draw() {
	source := rl.NewRectangle(0, 0, float32(s.Texture.Width), float32(s.Texture.Height))
	dest := rl.NewRectangle(s.Pos.X, s.Pos.Y, s.Width, s.Height)
	origin := rl.NewVector2(s.Width/2, s.Height)

	rl.DrawTexturePro(s.Texture, source, dest, origin, s.rotation, rl.White)
}
