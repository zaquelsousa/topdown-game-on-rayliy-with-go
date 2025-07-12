package maps

import (
	"encoding/json"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadMap() {

	TileSet := InitTiles()

	matrixData, err := os.ReadFile("maps/matrixmap.json")

	var matrixMap [][]int
	err = json.Unmarshal(matrixData, &matrixMap)

	if err != nil {
		fmt.Println("deu merda: ", err)
	}

	tileSize := int32(16)
	for y, row := range matrixMap {
		for x, TileID := range row {
			PosX := int32(x) * tileSize
			PosY := int32(y) * tileSize

			switch TileSet[TileID] {
			case TileSet[1]:
				rl.DrawTexture(TileSet[1].Texture, PosX, PosY, rl.White)

			case TileSet[2]:
				rl.DrawTexture(TileSet[2].Texture, PosX, PosY, rl.White)
			case TileSet[3]:
				rl.DrawTexture(TileSet[3].Texture, PosX, PosY, rl.White)

			default:
				rl.DrawRectangle(PosX, PosY, tileSize, tileSize, rl.Black)
			}
		}
	}
}
