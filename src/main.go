package main

import rl "github.com/gen2brain/raylib-go/raylib"
import rg "github.com/gen2brain/raylib-go/raygui"

var camera rl.Camera3D = rl.NewCamera3D(
	rl.NewVector3(10, 10, 10),
	rl.NewVector3(0, 0, 0),
	rl.NewVector3(0, 1, 0),
	60,
	rl.CameraPerspective,
)

func main() {

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(800, 450, "Mineclone")
	defer rl.CloseWindow()
	rl.DisableCursor()
	rl.SetTargetFPS(60)
	rl.ToggleFullscreen()

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraFree)

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawFPS(10, 10)

		rl.BeginMode3D(camera)
		rl.DrawGrid(10, 1)

		rl.EndMode3D()

		rl.EndDrawing()
	}
}
