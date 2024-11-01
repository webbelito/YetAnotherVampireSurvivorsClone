package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Window settings
const SCREEN_WIDTH = 1920
const SCREEN_HEIGHT = 1080
const SCREEN_TITLE = "YAVSC - Yet Another Vampire Survivors Clone"

const DEBUG_MODE = true

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

	rl.SetTargetFPS(180)

	game.Run()

}

// TODO: Implement a proper debug mode
func isDebugMode() bool {
	return DEBUG_MODE
}
