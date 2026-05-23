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
	var gamepad int32 = 0

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		if rl.IsGamepadAvailable(gamepad) {
			rl.DrawText(fmt.Sprintf("GP1: %s", rl.GetGamepadName(gamepad)), 10, 10, 40, rl.Black)
		}else {
			rl.DrawText("Controller not connected", 10, 10, 40, rl.Black)	
		}
		rl.EndDrawing()
	}
}