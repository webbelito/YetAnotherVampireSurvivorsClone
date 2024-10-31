package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const FixedUpdateRate float32 = 60.0

type Game struct {
	Player          *PlayerCharacter
	Level           *Level
	Camera          rl.Camera2D
	Enemies         []*Enemy
	Projectiles     []*Projectile
	PowerUps        []*PowerUp
	FixedDeltaTime  float32
	LastFixedUpdate time.Time
}

func NewGame() *Game {

	grid := []string{
		"@^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^@",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"@__________________________________________________________@",
	}

	level := NewLevel(grid)
	camera := rl.Camera2D{
		Offset: rl.Vector2{
			X: float32(rl.GetScreenWidth()) / 2,
			Y: float32(rl.GetScreenHeight()) / 2,
		},
		Target: rl.Vector2{
			X: float32(rl.GetScreenWidth()) / 2,
			Y: float32(rl.GetScreenHeight()) / 2,
		},
		Rotation: 0,
		Zoom:     2,
	}

	return &Game{
		Player:          nil,
		Level:           level,
		Camera:          camera,
		Enemies:         make([]*Enemy, 0),
		Projectiles:     make([]*Projectile, 0),
		PowerUps:        make([]*PowerUp, 0),
		FixedDeltaTime:  1.0 / FixedUpdateRate,
		LastFixedUpdate: time.Now(),
	}
}

// TODO: Remove this with a proper way to handle frame updates
var fixedFrameCounter int

func (g *Game) Update() {

	if g.Player != nil && g.Player.IsDead {
		rl.DrawText("Game Over", int32(rl.GetScreenWidth())/2-250, int32(rl.GetScreenHeight())/2, 100, rl.Red)
		return
	}

	// TODO: Create a better way of spawning the player
	if g.Player == nil {
		g.SpawnPlayer()
	}

	// Initialize the HUD
	if g.Player.HUD == nil {
		g.Player.HUD = NewHUD(g.Player)
	}

	// Spawn an enemy
	if rl.IsKeyPressed(rl.KeyB) {

		for i := 0; i < 24; i++ {
			g.SpawnBat()
		}

		g.SpawnBat()
	}

	// Spawn a pumpkin
	if rl.IsKeyPressed(rl.KeySpace) {
		g.SpawnPumpkin()
	}

	// Spawn a powerup
	if rl.IsKeyPressed(rl.KeyP) {
		g.SpawnPowerUp()
	}

	// Player takes damage
	if rl.IsKeyPressed(rl.KeyK) {
		g.Player.TakeDamage(10)
	}

	// Player heals
	if rl.IsKeyPressed(rl.KeyH) {
		g.Player.Heal(10)
	}

	// Player gains experience
	if rl.IsKeyPressed(rl.KeyE) {
		g.Player.GainExperience(10)
	}

	// Update player
	g.Player.Update(g)

	// Update enemies
	for i := 0; i < len(g.Enemies); i++ {

		if g.Enemies[i].IsDead {
			continue
		}

		g.Enemies[i].Update(g.Player)
	}

	// Update projectiles
	for i := 0; i < len(g.Projectiles); i++ {

		if !g.Projectiles[i].Active {
			continue
		}

		g.Projectiles[i].Update(g)
	}

	// Update powerups
	for i := 0; i < len(g.PowerUps); i++ {
		g.PowerUps[i].Update(g)
	}

	g.DestroyProjectiles()
	g.DestroyEnemy()
	g.DestroyPowerUp()

	// Update the camera
	g.Camera.Target = rl.Vector2{
		X: g.Player.X,
		Y: g.Player.Y,
	}

}

func (g *Game) FixedUpdate() {
	fixedFrameCounter++

	// Player
	g.Player.FixedUpdate(g)

	// Resolve collisions every 3rd frame
	if fixedFrameCounter%3 == 0 {
		ResolveEnemyCollisions(g)
	}

	// Enemies
	for i := 0; i < len(g.Enemies); i++ {
		g.Enemies[i].FixedUpdate(g)
	}
}

func (g *Game) Render() {
	// Level stuffs here
	g.Level.Render()

	// Player stuffs here
	g.Player.UpdateAnimation()
	g.Player.Render()

	// Enemy stuffs here
	for _, e := range g.Enemies {
		e.UpdateAnimation()
		e.Render()
	}

	// Projectile stuffs here
	for _, p := range g.Projectiles {
		p.Render()
	}
}

func (g *Game) Run() {
	for !rl.WindowShouldClose() {

		now := time.Now()
		deltaTime := time.Since(g.LastFixedUpdate).Seconds()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(g.Camera)

		g.Update()

		for deltaTime >= float64(g.FixedDeltaTime) {
			g.FixedUpdate()
			g.LastFixedUpdate = now
			deltaTime -= float64(g.FixedDeltaTime)
		}

		g.Render()

		rl.EndMode2D()

		// HUD Stuffs Here
		g.Player.HUD.Render()
		g.RenderMobsCounter()
		g.RenderPowerUpHUD()

		// Draw the FPS
		rl.DrawFPS(10, int32(rl.GetScreenHeight()-20))

		rl.EndDrawing()
	}
}

func (g *Game) SpawnPlayer() {
	g.Player = NewPlayer("Bill", 32, 64, 150, 100, 50)
}

func (g *Game) SpawnProjectile(x, y, radius, speed float32, direction rl.Vector2, color rl.Color, isHoming bool) {
	g.Projectiles = append(g.Projectiles, NewProjectile(TextureAtlas, x, y, radius, speed, direction, color, isHoming))
}

func (g *Game) DestroyProjectiles() {

	if len(g.Projectiles) == 0 {
		return
	}

	i := 0

	for _, p := range g.Projectiles {
		if p.Active {
			g.Projectiles[i] = p
			i++
		}
	}

	// Truncate the slice to remove inactive projectiles
	g.Projectiles = g.Projectiles[:i]
}

func (g *Game) SpawnBat() {
	g.Enemies = append(g.Enemies, NewEnemy(TextureAtlas, "Bat", 32, 32, 10))
}

func (g *Game) SpawnPumpkin() {
	g.Enemies = append(g.Enemies, NewEnemy(TextureAtlas, "Pumpkin", 32, 32, 20))
}

func (g *Game) DestroyEnemy() {

	if len(g.Enemies) == 0 {
		return
	}

	i := 0

	for _, e := range g.Enemies {
		if !e.IsDead {
			g.Enemies[i] = e
			i++
		}
	}

	// Truncate the slice to remove dead enemies
	g.Enemies = g.Enemies[:i]

}

func (g *Game) SpawnPowerUp() {

	// Build a new powerup
	powerUp := NewPowerUp()

	// Randomize the PowerUpType
	powerUp.RandomizePowerUpType()

	// Randomize the spawn point
	powerUp.Position = powerUp.RandomizeSpawnPoint()

	// Append the new powerup to the game's powerups
	g.PowerUps = append(g.PowerUps, powerUp)

}

// TODO: Replace temporary function to Render PowerUp HUD
func (g *Game) RenderPowerUpHUD() {
	for i, pu := range g.Player.PowerUps {
		pu.RenderHUD(20 * int32(i))
	}
}

func (g *Game) RenderMobsCounter() {
	mobCount := fmt.Sprintf("Mobs: %d", len(g.Enemies))
	rl.DrawText(mobCount, 10, int32(rl.GetScreenHeight()-50), 20, rl.Red)
}

func (g *Game) DestroyPowerUp() {

	if len(g.PowerUps) == 0 {
		return
	}

	// Remove the powerup from the game if its is not active and has been picked up
	i := 0

	for _, pu := range g.PowerUps {
		if !pu.Expired {
			g.PowerUps[i] = pu
			i++
		}
	}

	// Truncate the slice to remove inactive powerups
	g.PowerUps = g.PowerUps[:i]

}
