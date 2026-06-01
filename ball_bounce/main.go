package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	maxBuildings int = 100
)

func main() {
	fmt.Println("HelloWorld")
	screenWidth := int32(800)
	screenHeight := int32(450)

	// rl.SetConfigFlags(rl.FlagMsaa4xHint)

	rl.InitWindow(screenWidth, screenHeight, "Ball Bounce")
	defer rl.CloseWindow()

	ballPosition := rl.Vector2{X: float32(screenWidth) / 2, Y: float32(screenHeight) / 2}

	paused := false
	useGravity := true

	ballSpeed := rl.Vector2{X: 5.0, Y: 4.0}
	ballRadius := 30
	gravity := 0.8
	// energyLoss := 0.01

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeySpace) {
			paused = !paused
		}
		if rl.IsKeyPressed(rl.KeyG) {
			useGravity = !useGravity
		}

		if !paused {
			ballPosition.X += ballSpeed.X
			ballPosition.Y += ballSpeed.Y

			if useGravity {
				// ballSpeed.X = ballSpeed.X - ballSpeed.X*float32(energyLoss)
				ballSpeed.Y += float32(gravity)
			}

			if (ballPosition.X <= float32(ballRadius)) || (ballPosition.X >= float32(screenWidth)-float32(ballRadius)) {
				ballSpeed.X *= -1
			}
			if (ballPosition.Y <= float32(ballRadius)) || (ballPosition.Y >= (float32(screenHeight) - float32(ballRadius))) {
				ballSpeed.Y *= -1
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rl.DrawCircleV(ballPosition, float32(ballRadius), rl.Red)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

}
