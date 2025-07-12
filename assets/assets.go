package assets

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	PlayerSprite rl.Texture2D
	SwordSprite  rl.Texture2D
	GrassSprite  rl.Texture2D
	RockSprite   rl.Texture2D
	WaterSprite  rl.Texture2D
)

func LoadTextures() {
	PlayerSprite = rl.LoadTexture("player/sprites/player.png")
	SwordSprite = rl.LoadTexture("weapons/sprites/sword.png")
	GrassSprite = rl.LoadTexture("maps/sprites/grass.png")
	RockSprite = rl.LoadTexture("maps/sprites/rock.png")
	WaterSprite = rl.LoadTexture("maps/sprites/water.png")

}

func UnloadTextures() {
	rl.UnloadTexture(PlayerSprite)
	rl.UnloadTexture(SwordSprite)
	rl.UnloadTexture(GrassSprite)
	rl.UnloadTexture(RockSprite)
	rl.UnloadTexture(WaterSprite)

}
