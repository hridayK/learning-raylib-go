// new code:

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Realistic Rectangle Collision")
	defer rl.CloseWindow()

	rec1 := rl.Rectangle{X: 225, Y: 225, Width: 160, Height: 60}
	rec2 := rl.Rectangle{X: 125, Y: 125, Width: 80, Height: 80}

	// Speed and movement direction for rec1
	speed := float32(200.0)
	direction := float32(1.0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// 1. Update rec2 to follow mouse center (feels better than top-left corner)
		mousePos := rl.GetMousePosition()
		rec2.X = mousePos.X - (rec2.Width / 2)
		rec2.Y = mousePos.Y - (rec2.Height / 2)

		// 2. Move rec1
		rec1.X += speed * direction * rl.GetFrameTime()

		// 3. Screen boundary collision (Bounce off walls)
		if rec1.X <= 0 {
			rec1.X = 0 // Reset position to prevent getting stuck
			direction = 1.0
		} else if (rec1.X + rec1.Width) >= float32(rl.GetScreenWidth()) {
			rec1.X = float32(rl.GetScreenWidth()) - rec1.Width
			direction = -1.0
		}

		overlap := rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0}

		// 4. Realistic Box-to-Box Collision Response
		if rl.CheckCollisionRecs(rec1, rec2) {
			// Get the exact overlap rectangle to see where they intersected
			overlap = rl.GetCollisionRec(rec1, rec2)

			// Side collision check: horizontal overlap is smaller than vertical overlap
			if overlap.Width < overlap.Height {

				// CONDITION A: Block is moving RIGHT, and hits the LEFT side of the mouse box
				if direction > 0 && rec1.X < rec2.X {
					rec1.X = rec2.X - rec1.Width // Snap to left edge
					direction = -1.0             // Bounce left
				}

				// CONDITION B: Block is moving LEFT, and hits the RIGHT side of the mouse box
				if direction < 0 && rec1.X > rec2.X {
					rec1.X = rec2.X + rec2.Width // Snap to right edge
					direction = 1.0              // Bounce right
				}

				// If neither condition is met (e.g., block is moving right, but the mouse
				// clipped it from behind), the block ignores the collision and keeps going!
			}
		}

		// 5. Drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Change color dynamically on collision
		rec1Color := rl.Red
		// if rl.CheckCollisionRecs(rec1, rec2) {
		// 	rec1Color = rl.Orange
		// }

		rl.DrawRectangleRec(rec1, rec1Color)
		rl.DrawRectangleRec(rec2, rl.Fade(rl.Blue, 0.8)) // Slight transparency for the mouse box
		rl.DrawRectangle(int32(overlap.X), int32(overlap.Y), int32(overlap.Width), int32(overlap.Height), rl.Violet)
		rl.DrawText("Block will bounce off the mouse rectangle!", 10, 10, 20, rl.DarkGray)
		rl.EndDrawing()
	}
}

// old code:

// package main

// import (
// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

// func main() {
// 	rl.InitWindow(800, 450, "Basic Title")
// 	defer rl.CloseWindow()

// 	rec1 := rl.Rectangle{X: 225, Y: 225, Width: 160, Height: 60}
// 	rec2 := rl.Rectangle{X: 125, Y: 125, Width: 80, Height: 80}
// 	paused := false
// 	direction := 1

// 	collision := false

// 	for !rl.WindowShouldClose() {
// 		rec2.X = rl.GetMousePosition().X
// 		rec2.Y = rl.GetMousePosition().Y
// 		if !paused {
// 			rec1.X += float32(100*direction) * rl.GetFrameTime()
// 		}

// 		collision = rl.CheckCollisionRecs(rec1, rec2)

// 		if rl.IsMouseButtonDown(rl.MouseLeftButton) || collision {
// 			paused = !paused
// 			collision = rl.CheckCollisionRecs(rec1, rec2)
// 		}

// 		if (rec1.X <= 0) || ((rec1.X + rec1.Width) >= float32(rl.GetScreenWidth())) {
// 			direction *= -1
// 		}

// 		rl.BeginDrawing()
// 		rl.ClearBackground(rl.White)
// 		rl.DrawRectangle(int32(rec1.X), int32(rec1.Y), int32(rec1.Width), int32(rec1.Height), rl.Red)
// 		rl.DrawRectangle(int32(rec2.X), int32(rec2.Y), int32(rec2.Width), int32(rec2.Height), rl.Blue)
// 		rl.EndDrawing()
// 	}
// }
