package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const SCREEN_WIDTH = 1920
const SCREEN_HEIGHT = 1080
const SCREEN_TITLE = "YAVSC - Yet Another Vampire Survivors Clone"

func main() {

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)

	rl.SetTargetFPS(144)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(SCREEN_TITLE, SCREEN_WIDTH/2-500, SCREEN_HEIGHT/2, 48, rl.Maroon)

		rl.EndDrawing()
	}
}
