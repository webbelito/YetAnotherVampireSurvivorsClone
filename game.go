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

	// TODO: Create a better way of spawning the player
	if g.Player == nil {
		g.SpawnPlayer()
	}

	// Spawn an enemy
	if rl.IsKeyPressed(rl.KeySpace) {
		g.SpawnEnemy()
	}

	// Spawn a powerup
	if rl.IsKeyPressed(rl.KeyP) {
		g.SpawnPowerUp()
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

	g.DestroyProjectiles()
	g.DestroyEnemy()
	g.DestroyPowerUp()

}

func (g *Game) SpawnPlayer() {
	g.Player = NewPlayer("Dracula", 50, 100, 150, 100, 50)
}

func (g *Game) SpawnProjectile(x, y, radius, speed float32, direction rl.Vector2, color rl.Color, isHoming bool) {
	g.Projectiles = append(g.Projectiles, NewProjectile(x, y, radius, speed, direction, color, isHoming))
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

func (g *Game) SpawnEnemy() {
	g.Enemies = append(g.Enemies, NewEnemy("Bat", 50, 50, 100, 10, 50))
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
