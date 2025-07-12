package enemys

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/assets"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/weapons"
)

type LongRange struct {
	Life         float32
	Position     rl.Vector2
	Sprite       rl.Texture2D
	FrameRec     rl.Rectangle
	Speed        float32
	Color        rl.Color
	width        float32
	height       float32
	WasHited     bool
	FleeCooldown float32
}

func NewLongRange() LongRange {
	return LongRange{
		Life:     100,
		Position: rl.NewVector2(159, 112),
		Sprite:   assets.PlayerSprite,
		Speed:    200,
		Color:    rl.Blue,
		width:    16,
		height:   32,
		WasHited: false,
	}
}

func (l *LongRange) GotHit(ataque weapons.Sword) {

	var hitRange = float32(30)
	dist := rl.Vector2Distance(l.Position, ataque.Pos)

	if dist <= hitRange && ataque.IsAttk && !l.WasHited {
		fmt.Println("hit")
		l.Life -= float32(ataque.Damege)
		l.WasHited = true
		l.FleeCooldown = 2.0
	}
	if !ataque.IsAttk {
		l.WasHited = false
	}
}

func (l *LongRange) Update(playerPos rl.Vector2, ataque weapons.Sword, others []LongRange) {
	if l.Life > 0 {
		if l.FleeCooldown > 0 {
			l.FleeCooldown -= rl.GetFrameTime()
		}

		if rl.Vector2Distance(l.Position, playerPos) <= 100 && l.FleeCooldown > 0 {
			fmt.Println("get away from me")
			direction := rl.Vector2Subtract(l.Position, playerPos)
			direction = rl.Vector2Normalize(direction)

			newPos := rl.Vector2Add(playerPos, rl.Vector2Scale(direction, 100))
			l.Position = rl.Vector2MoveTowards(l.Position, newPos, l.Speed*rl.GetFrameTime())
		}

		separationDistance := float32(20)
		moveAway := rl.NewVector2(0, 0)

		for _, other := range others {
			if &other == l || other.Life <= 0 {
				continue
			}

			dist := rl.Vector2Distance(l.Position, other.Position)
			if dist > 0 && dist < separationDistance {
				diff := rl.Vector2Subtract(l.Position, other.Position)
				diff = rl.Vector2Normalize(diff)
				diff = rl.Vector2Scale(diff, separationDistance-dist)
				moveAway = rl.Vector2Add(moveAway, diff)
			}
		}

		l.Position = rl.Vector2Add(l.Position, moveAway)
		l.GotHit(ataque)
	}
}

func (e *LongRange) Draw() {
	if e.Life > 0 {
		maxBarWidth := float32(50)
		lifePercent := e.Life / 100
		barWidth := int32(maxBarWidth * lifePercent)
		barX := int32(e.Position.X + e.width/2 - float32(barWidth)/2)
		barY := int32(e.Position.Y) - 10
		rl.DrawRectangle(barX, barY, barWidth, 5, rl.Red)
		rl.DrawRectangle(int32(e.Position.X), int32(e.Position.Y), int32(e.width), int32(e.height), rl.Purple)
	}
}
