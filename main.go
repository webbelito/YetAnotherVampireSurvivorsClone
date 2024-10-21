package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/enemy"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/player"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/projectile"
)

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

	player := player.NewPlayer("Dracula", PLAYER_WIDTH, PLAYER_HEIGHT, PLAYER_SPEED, PLAYER_HEALTH, PLAYER_DAMAGE)

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)

	rl.SetTargetFPS(144)
	fmt.Println(player.GetName())

	// TODO: Implement a map of enemies
	enemy.NewEnemy("Bat", BAT_WIDTH, BAT_HEIGHT, BAT_SPEED, BAT_HEALTH, BAT_DAMAGE)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// TODO: Implement a better way to spawn enemies
		if rl.IsKeyPressed(rl.KeySpace) {
			enemy.NewEnemy("Bat", BAT_WIDTH, BAT_HEIGHT, BAT_SPEED, BAT_HEALTH, BAT_DAMAGE)
		}

		if rl.IsKeyReleased(rl.KeyT) {
		}

		if rl.IsKeyReleased(rl.KeyR) {
		}

		player.Update()
		enemy.Update(player)

		projectile.Update()

		rl.DrawText(SCREEN_TITLE, SCREEN_WIDTH/2-500, SCREEN_HEIGHT/2, 48, rl.Maroon)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
