package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Window settings
const SCREEN_WIDTH = 1920
const SCREEN_HEIGHT = 1080
const SCREEN_TITLE = "YAVSC - Yet Another Vampire Survivors Clone"

var TextureAtlas rl.Texture2D

func main() {

	// Check if we are in debug mode and set the log level
	if isDebugMode() {
		rl.SetTraceLogLevel(rl.LogDebug)
	}

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)
	defer rl.CloseWindow()

	// Load the atlas texture
	TextureAtlas = rl.LoadTexture("assets/images/yavsc-atlas-sheet.png")
	defer rl.UnloadTexture(TextureAtlas)

	// Check that the atlast loaded correctly
	if TextureAtlas.Width == 0 || TextureAtlas.Height == 0 {
		rl.TraceLog(rl.LogError, "Failed to load the atlas texture")
		return
	}

	// Initialize the game
	game := NewGame()

	// Random timer to despawn title text
	titleDisplayTimer := float32(0.0)

	rl.SetTargetFPS(180)

	for !rl.WindowShouldClose() {

		// TODO: Implement a 2d camera

		rl.BeginDrawing()

		game.Update()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawFPS(int32(10), int32(rl.GetScreenHeight()-30))

		// Display title text for 3 seconds
		if titleDisplayTimer < 3.0 {
			titleDisplayTimer = titleDisplayTimer + rl.GetFrameTime()
			rl.DrawText(SCREEN_TITLE, SCREEN_WIDTH/2-500, SCREEN_HEIGHT/2, 48, rl.Maroon)
		}

		rl.EndDrawing()
	}
}

// TODO: Implement a proper debug mode
func isDebugMode() bool {
	debug := true
	return debug
}
