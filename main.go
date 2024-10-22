package main

import ( /*"github.com/webbelito/YetAnotherVampireSurvivorsClone/enemy"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/player"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/projectile"
	*/ // 3rd-party library
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/game"
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

	// Initialize the game
	game := game.NewGame()

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)

	rl.SetTargetFPS(180)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		game.Update()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawFPS(10, 10)

		rl.DrawText(SCREEN_TITLE, SCREEN_WIDTH/2-500, SCREEN_HEIGHT/2, 48, rl.Maroon)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

/*
func InitGame() {

	// Initialize the Game struct
	game := &Game{}

	game.playerObj = player.NewPlayer("Dracula", PLAYER_WIDTH, PLAYER_HEIGHT, PLAYER_SPEED, PLAYER_HEALTH, PLAYER_DAMAGE)
	game.enemies = []*enemy.Enemy{}
	game.projectiles = []*projectile.Projectile{}

}

func UpdateGame() {

	game.playerObj.Update()

	// Update all enemies
	for i := 0; i < len(game.enemies); i++ {
		game.enemies[i].Update(game.playerObj)
	}

	// Update all projectiles
	for i := 0; i < len(game.projectiles); i++ {
		game.projectiles[i].Update()
	}

	// Spawn 'BAT' enemies
	if rl.IsKeyPressed(rl.KeySpace) {
		game.enemies = append(game.enemies, enemy.NewEnemy("BAT", BAT_WIDTH, BAT_HEIGHT, BAT_SPEED, BAT_HEALTH, BAT_DAMAGE))
	}
}

func (g *Game) SpawnProjectile(x float32, y float32, radius float32, speed float32, direction rl.Vector2) {
	g.projectiles = append(g.projectiles, projectile.NewProjectile(x, y, radius, speed, direction))

}
*/
