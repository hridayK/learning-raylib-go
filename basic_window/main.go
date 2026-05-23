package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main(){
	rl.InitWindow(800,450, "Basic Window's title text")
	defer rl.CloseWindow() // defer means it will execute at the end of the function.
	// defer is used for clean up functions.
	// it ensure that rl.CloseWindow function runs even if a runtime error causes an early exit.

	for !rl.WindowShouldClose(){
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Hello",200,200,22,rl.Black)
		rl.EndDrawing()
	}
}