package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/enemy"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/player"
)

const SCREEN_WIDTH = 1920
const SCREEN_HEIGHT = 1080
const SCREEN_TITLE = "YAVSC - Yet Another Vampire Survivors Clone"

func main() {

	player := player.NewPlayer("Dracula", 50, 100, 150)

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)

	rl.SetTargetFPS(60)
	fmt.Println(player.GetName())

	enemy := enemy.NewEnemy("Zombie", 50, 100)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyPressed(rl.KeySpace) {
			enemy.Spawn()
		}

		player.Update()
		enemy.Update()

		rl.DrawText(SCREEN_TITLE, SCREEN_WIDTH/2-500, SCREEN_HEIGHT/2, 48, rl.Maroon)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
