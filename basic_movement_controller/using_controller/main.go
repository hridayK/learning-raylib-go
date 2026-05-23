package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	xboxNameId="Xbox Controller"
	playstaionNameId="Playstation Controller"
)

func main() {

	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	screenWidth := int32(800)
	screenHeight := int32(450)


	rl.InitWindow(screenWidth,screenHeight, "Basic Movement - Controller")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ballPosition := rl.Vector2{X: 80, Y: 80}
	var gamepad int32 = 0

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawCircle(int32(ballPosition.X), int32(ballPosition.Y), 30, rl.Red)
		if rl.IsGamepadAvailable(gamepad) {
			rl.DrawText(fmt.Sprintf("GP1: %s", rl.GetGamepadName(gamepad)), 10, 10, 20, rl.Black)
			
			if rl.IsGamepadButtonDown(gamepad, rl.GamepadButtonLeftFaceDown) {
				ballPosition.Y += 5
			}
			if rl.IsGamepadButtonDown(gamepad, rl.GamepadButtonLeftFaceUp) {
				ballPosition.Y -= 5
			}
			if rl.IsGamepadButtonDown(gamepad, rl.GamepadButtonLeftFaceRight) {
				ballPosition.X += 5
			}
			if rl.IsGamepadButtonDown(gamepad, rl.GamepadButtonLeftFaceLeft) {
				ballPosition.X -= 5
			}

			ballPosition.X += rl.GetGamepadAxisMovement(gamepad, rl.GamepadAxisLeftX) * 5
			ballPosition.Y += rl.GetGamepadAxisMovement(gamepad, rl.GamepadAxisLeftY) * 5

		}else {
			rl.DrawText("Controller not connected", 10, 10, 20, rl.Black)	
		}
		rl.EndDrawing()
	}
}