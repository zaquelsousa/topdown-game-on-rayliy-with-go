package maps

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/assets"
)

type Tile struct {
	Texture      rl.Texture2D
	Size         float32
	IsCollidable bool
}

func InitTiles() map[int]Tile {

	var Grass = Tile{
		Texture:      assets.GrassSprite,
		Size:         16,
		IsCollidable: false,
	}

	var Wall = Tile{
		Texture:      assets.RockSprite,
		Size:         16,
		IsCollidable: true,
	}
	var Water = Tile{
		Texture:      assets.WaterSprite,
		Size:         16,
		IsCollidable: true,
	}
	var TileSet = map[int]Tile{
		1: Grass,
		2: Wall,
		3: Water,
	}

	return TileSet
}
