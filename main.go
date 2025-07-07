package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/player"
)

// game congif
const (
	screenWidth  = int32(512)
	screenHeight = int32(224)
)

type Obj struct {
	Pos    rl.Vector2
	Width  float32
	Height float32
}

// Game struct
type Game struct {
	Player player.Player
	Camera rl.Camera2D
	obj    Obj
}

// Initialize the game state
func (g *Game) Init() {
	plrSprite := rl.LoadTexture("player/sprites/player.png")
	g.Player = player.NewPlayer(plrSprite)

	//camera configs
	g.Camera.Target = rl.NewVector2(g.Player.Position.X+20, g.Player.Position.Y+20)
	g.Camera.Offset = rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	g.Camera.Rotation = 0.0
	g.Camera.Zoom = 2.0

	g.obj = Obj{
		Pos:    rl.NewVector2(30, float32(screenHeight/2-50)),
		Width:  10,
		Height: 100,
	}
}

// Game update logic
func (g *Game) Update() {
	g.Player.Update()
	g.Camera.Target = rl.NewVector2(g.Player.Position.X+20, g.Player.Position.Y+20)
}

// Draw game elements
func (g *Game) Draw() {
	rl.BeginMode2D(g.Camera)
	g.Player.Draw()
	rl.DrawRectangleV(g.obj.Pos, rl.NewVector2(g.obj.Width, g.obj.Height), rl.Black)
	rl.EndMode2D()
}

func main() {
	screenWidth := int32(512)
	screenHeight := int32(224)

	rl.InitWindow(screenWidth, screenHeight, "raylib -go template")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var game Game
	game.Init()

	// Carregar shader
	shader := rl.LoadShader("", "bkg.fs")
	defer rl.UnloadShader(shader)

	// Pegar location de uniforms
	resLoc := rl.GetShaderLocation(shader, "resolution")

	for !rl.WindowShouldClose() {

		game.Update()
		// Atualiza resolução
		res := []float32{float32(screenWidth), float32(screenHeight)}
		rl.SetShaderValue(shader, resLoc, res, rl.ShaderUniformVec2)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginShaderMode(shader)
		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Green)
		rl.EndShaderMode()
		game.Draw()
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
