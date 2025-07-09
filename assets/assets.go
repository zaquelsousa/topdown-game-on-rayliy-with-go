package assets

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	PlayerSprite rl.Texture2D
	SwordSprite  rl.Texture2D
	// outras texturas que for usar...
)

func LoadTextures() {
	PlayerSprite = rl.LoadTexture("player/sprites/player.png")
	SwordSprite = rl.LoadTexture("weapons/sprites/sword.png")
	// carrega outras aqui
}

func UnloadTextures() {
	rl.UnloadTexture(PlayerSprite)
	rl.UnloadTexture(SwordSprite)
	// descarrega outras aqui
}
