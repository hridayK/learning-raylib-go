package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "This is the title")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ballPosition := rl.Vector2{X: 80, Y: 80}

	for !rl.WindowShouldClose() {

		ballPosition = rl.GetMousePosition()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawCircle(int32(ballPosition.X), int32(ballPosition.Y), 50, rl.Red)
		rl.EndDrawing()
	}
}
