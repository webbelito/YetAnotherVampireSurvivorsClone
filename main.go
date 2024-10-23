package main

import ( /*"github.com/webbelito/YetAnotherVampireSurvivorsClone/enemy"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/player"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/projectile"
	*/ // 3rd-party library
	rl "github.com/gen2brain/raylib-go/raylib"
	/*
		"github.com/webbelito/YetAnotherVampireSurvivorsClone/game"
	*/)

// Window settings
const SCREEN_WIDTH = 1920
const SCREEN_HEIGHT = 1080
const SCREEN_TITLE = "YAVSC - Yet Another Vampire Survivors Clone"

// Player settings
const PLAYER_WIDTH = 50
const PLAYER_HEIGHT = 100
const PLAYER_SPEED = 150
const PLAYER_HEALTH = 100
const PLAYER_DAMAGE = 50

// Enemy settings
const BAT_WIDTH = 50
const BAT_HEIGHT = 50
const BAT_SPEED = 50
const BAT_HEALTH = 50
const BAT_DAMAGE = 10

func main() {

	// Initialize the game
	game := NewGame()

	// Random timer to despawn title text
	titleDisplayTimer := float32(0.0)

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)

	rl.SetTargetFPS(180)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		game.Update()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawFPS(10, 10)

		if titleDisplayTimer < 3.0 {
			titleDisplayTimer = titleDisplayTimer + rl.GetFrameTime()
			rl.DrawText(SCREEN_TITLE, SCREEN_WIDTH/2-500, SCREEN_HEIGHT/2, 48, rl.Maroon)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
