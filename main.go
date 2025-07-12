package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/assets"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/maps"
	"github.com/zaquelsousa/topdown-game-on-rayliy-with-go/player"
)

// game congif
const (
	screenWidth  = int32(512)
	screenHeight = int32(224)
)

// Game struct
type Game struct {
	Player      player.Player
	Camera      rl.Camera2D
	CurrentRoom maps.Room
}

// Initialize the game state
func (g *Game) Init() {
	g.Player = player.NewPlayer()
	g.CurrentRoom = maps.NewRoom()
	g.CurrentRoom.SpawnEnemy()

	//camera configs
	g.Camera.Target = rl.NewVector2(g.Player.Position.X+20, g.Player.Position.Y+20)
	g.Camera.Offset = rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	g.Camera.Rotation = 0.0
	g.Camera.Zoom = 1.5
}

// Game update logic
func (g *Game) Update() {
	g.Player.Update(g.Camera)
	g.CurrentRoom.Update(g.Player.Position, g.Player.Sword)
	g.Camera.Target = rl.NewVector2(g.Player.Position.X+20, g.Player.Position.Y+20)
}

// Draw game elements
func (g *Game) Draw() {
	rl.BeginMode2D(g.Camera)
	maps.LoadMap()
	g.CurrentRoom.Draw()
	g.Player.Draw()
	rl.EndMode2D()
}

func main() {
	//flg to resize the window
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(screenWidth, screenHeight, "raylib -go template")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	assets.LoadTextures()
	defer assets.UnloadTextures()
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
