package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Player      *PlayerCharacter
	Enemies     []*Enemy
	Projectiles []*Projectile
	PowerUps    []*PowerUp
}

func NewGame() *Game {
	return &Game{
		Enemies:     make([]*Enemy, 0),
		Projectiles: make([]*Projectile, 0),
		PowerUps:    make([]*PowerUp, 0),
	}
}

func (g *Game) Update() {

	if g.Player != nil && g.Player.IsDead {
		rl.DrawText("Game Over", int32(rl.GetScreenWidth())/2-250, int32(rl.GetScreenHeight())/2, 100, rl.Red)
		return
	}

	// Initialize the atlas
	/*
		if g.Atlas.ID == 0 {

			// Load the atlas texture
			g.Atlas = rl.LoadTexture("assets/images/yavsc-atlast-sheet.png")

			// Check that the atlast loaded correctly
			if g.Atlas.Width == 0 || g.Atlas.Height == 0 {
				rl.TraceLog(rl.LogError, "Failed to load the atlas texture")
			}
		}
	*/

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

	// TODO: Render the PowerUp HUD
	g.Player.HUD.Render()

	g.RenderPowerUpHUD()

	g.DestroyProjectiles()
	g.DestroyEnemy()
	g.DestroyPowerUp()

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
