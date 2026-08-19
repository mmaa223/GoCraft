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

func ui() {
	rg.Button(rl.NewRectangle(10, 10, 100, 100), "hello")
}

func main() {

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(800, 450, "Mineclone")
	defer rl.CloseWindow()
	rl.DisableCursor()
	rl.SetTargetFPS(60)
	rl.ToggleFullscreen()

	var mesh rl.Mesh = rl.GenMeshCube(1, 1, 1)
	var checked rl.Image = *rl.GenImageChecked(2, 2, 1, 1, rl.Green, rl.Red)
	var texture rl.Texture2D = rl.LoadTextureFromImage(&checked)
	//var material rl.Material = rl.LoadMaterialDefault()
	var model rl.Model = rl.LoadModelFromMesh(mesh)

	rl.UnloadImage(&checked)

	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, texture)
	for !rl.WindowShouldClose() {
		//ui()
		rl.UpdateCamera(&camera, rl.CameraFree)

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawFPS(10, 10)

		rl.BeginMode3D(camera)
		rl.DrawGrid(10, 1)
		rl.DrawModel(model, rl.NewVector3(0, 0, 0), 1, rl.White)

		rl.EndMode3D()

		rl.EndDrawing()

	}
	rl.UnloadMesh(&mesh)
	rl.UnloadModel(model)

}
