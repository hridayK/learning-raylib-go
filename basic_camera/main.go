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

	rl.InitWindow(screenWidth, screenHeight, "This is the title")
	defer rl.CloseWindow()

	player := rl.NewRectangle(400, 280, 40, 40)

	buildings := make([]rl.Rectangle, maxBuildings)
	buildColors := make([]rl.Color, maxBuildings)

	spacing := float32(0)

	for i := 0; i < maxBuildings; i++ {
		r := rl.Rectangle{}

		r.Width = float32(rl.GetRandomValue(50, 200))
		r.Height = float32(rl.GetRandomValue(50, 200))

		r.Y = float32(screenHeight) - 130 - r.Height
		r.X = -6000 + spacing

		spacing += r.Width
		c := rl.NewColor(byte(rl.GetRandomValue(200, 240)), byte(rl.GetRandomValue(200, 240)), byte(rl.GetRandomValue(200, 240)), byte(255))

		buildings[i] = r
		buildColors[i] = c
	}

	camera := rl.Camera2D{}
	camera.Target = rl.NewVector2(float32(player.X)+20, float32(player.Y)+20)
	camera.Offset = rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2)
	camera.Rotation = 0.0
	camera.Zoom = 1.0

	rl.SetTargetFPS(60)

	// Declare keyName outside the game loop so it persists across frames
	keyName := ""

	for !rl.WindowShouldClose() {
		keyPressed := rl.GetKeyPressed()
		if keyPressed != 0 {
			if keyPressed == 263 {
				keyName = "Left Arrow Key"
			} else if keyPressed == 262 {
				keyName = "Right Arrow Key"
			} else {
				keyName = rl.GetKeyName(keyPressed)
			}
			fmt.Println(keyPressed)
		}

		if rl.IsKeyDown(rl.KeyRight) {
			player.X += 20
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			player.X -= 20
		}

		if rl.IsKeyDown(rl.KeyA) {
			camera.Rotation--
		} else if rl.IsKeyDown(rl.KeyS) {
			camera.Rotation++
		}

		zoomChange := rl.GetMouseWheelMove() * 0.05

		camera.Zoom += zoomChange

		if zoomChange != 0 {
			keyName = "Mouse Middle Wheel"
		}

		if camera.Zoom >= 3.0 {
			camera.Zoom = 3.0
		}
		if camera.Zoom <= 1.0 {
			camera.Zoom = 1.0
		}

		if rl.IsKeyPressed(rl.KeyR) {
			camera.Rotation = 0
			camera.Zoom = 1.0
		}

		if camera.Rotation > 40 {
			camera.Rotation = 40
		} else if camera.Rotation < -40 {
			camera.Rotation = -40
		}

		camera.Target = rl.NewVector2(float32(player.X+20), float32(player.Y+20))

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)

		rl.BeginMode2D(camera)

		rl.DrawRectangle(int32(camera.Target.X), -500, 1, screenHeight*4, rl.Green)
		rl.DrawRectangle(-500, int32(camera.Target.Y), screenWidth*4, 1, rl.Green)

		for i := range maxBuildings {
			rl.DrawRectangleRec(buildings[i], buildColors[i])
		}

		rl.DrawRectangleRec(player, rl.Red)
		rl.EndMode2D()

		// Draw Screen Borders
		rl.DrawRectangle(0, 0, screenWidth, 5, rl.Red)
		rl.DrawRectangle(0, 5, 5, screenHeight-10, rl.Red)
		rl.DrawRectangle(screenWidth-5, 5, 5, screenHeight-10, rl.Red)
		rl.DrawRectangle(0, screenHeight-5, screenWidth, 5, rl.Red)

		// UI Box dimensions
		boxX := int32(10)
		boxY := int32(10)
		boxWidth := int32(250)
		boxHeight := int32(113)
		fontSize := int32(20)

		// Draw UI Box Background and Border
		rl.DrawRectangle(boxX, boxY, boxWidth, boxHeight, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(boxX, boxY, boxWidth, boxHeight, rl.Blue)

		// Center and draw text dynamically
		if keyName != "" {
			textWidth := rl.MeasureText(keyName, fontSize)

			// Dynamic centering formulas
			textX := boxX + (boxWidth-textWidth)/2
			textY := boxY + (boxHeight-fontSize)/2

			rl.DrawText(keyName, textX, textY, fontSize, rl.Black)
		}

		rl.EndDrawing()
	}
}
