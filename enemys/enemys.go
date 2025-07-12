package enemys

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/assets"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/weapons"
)

type Enemy struct {
	Life     float32
	Position rl.Vector2
	Sprite   rl.Texture2D
	FrameRec rl.Rectangle
	Speed    float32
	Color    rl.Color
	width    float32
	height   float32
	WasHited bool
}

func NewEnemy() Enemy {
	return Enemy{
		Life:     100,
		Position: rl.NewVector2(159, 112),
		Sprite:   assets.PlayerSprite,
		Speed:    100,
		Color:    rl.Blue,
		width:    16,
		height:   32,
		WasHited: false,
	}
}

func (e *Enemy) GotHit(ataque weapons.Sword) {

	var hitRange = float32(30)
	dist := rl.Vector2Distance(e.Position, ataque.Pos)

	if dist <= hitRange && ataque.IsAttk && !e.WasHited {
		fmt.Println("hit")
		e.Life -= float32(ataque.Damege)
		e.WasHited = true
	}
	if !ataque.IsAttk {
		e.WasHited = false
	}
}

func (e *Enemy) Update(playerPos rl.Vector2, ataque weapons.Sword, others []Enemy) {
	if e.Life > 0 {
		e.Position = rl.Vector2MoveTowards(e.Position, playerPos, e.Speed*rl.GetFrameTime())

		separationDistance := float32(20)
		moveAway := rl.NewVector2(0, 0)

		for _, other := range others {
			if &other == e || other.Life <= 0 {
				continue
			}

			dist := rl.Vector2Distance(e.Position, other.Position)
			if dist > 0 && dist < separationDistance {
				diff := rl.Vector2Subtract(e.Position, other.Position)
				diff = rl.Vector2Normalize(diff)
				diff = rl.Vector2Scale(diff, separationDistance-dist)
				moveAway = rl.Vector2Add(moveAway, diff)
			}
		}

		e.Position = rl.Vector2Add(e.Position, moveAway)

		e.GotHit(ataque)
	}
}

func (e *Enemy) Draw() {
	if e.Life > 0 {
		rl.DrawRectangle(int32(e.Position.X), int32(e.Position.Y)-10, int32(e.Life), 5, rl.Red)
		rl.DrawRectangle(int32(e.Position.X), int32(e.Position.Y), int32(e.width), int32(e.height), rl.Brown)

	}
}
