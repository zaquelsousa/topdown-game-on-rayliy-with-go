package maps

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/enemys"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/weapons"
)

type Room struct {
	Enemies []enemys.LongRange
}

func NewRoom() Room {
	return Room{}
}
func (r *Room) SpawnEnemy() {
	enemyMaxNumber := 10

	for i := 0; i < enemyMaxNumber; i++ {
		l := enemys.NewLongRange()
		l.Position.X = float32(rl.GetRandomValue(10, 500))
		l.Position.Y = float32(rl.GetRandomValue(10, 200))

		r.Enemies = append(r.Enemies, l)
	}
}

func (r *Room) Update(playerPos rl.Vector2, ataque weapons.Sword) {
	for i := range r.Enemies {
		r.Enemies[i].Update(playerPos, ataque, r.Enemies)
	}
}

func (r *Room) Draw() {
	for i := range r.Enemies {
		r.Enemies[i].Draw()
	}
}
