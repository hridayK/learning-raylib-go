package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)


	rl.InitWindow(screenWidth,screenHeight, "Basic Movement")
	defer rl.CloseWindow()

	ballPosition := rl.Vector2{X: float32(screenWidth) / 2, Y: float32(screenHeight) / 2}
	ballRadius := float32(50)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
			ballPosition.X += 5
		}
		if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
			ballPosition.X -= 5
		}
		if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
			ballPosition.Y -= 5
		}
		if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
			ballPosition.Y += 5
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawCircle(int32(ballPosition.X), int32(ballPosition.Y), ballRadius, rl.Red)
		rl.EndDrawing()
	}
}