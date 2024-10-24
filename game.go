package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Player      *PlayerCharacter
	Enemies     []*Enemy
	Projectiles []*Projectile
}

func NewGame() *Game {
	return &Game{
		Enemies:     make([]*Enemy, 0),
		Projectiles: make([]*Projectile, 0),
	}
}

func (g *Game) Update() {

	if g.Player == nil {
		g.SpawnPlayer()
	}

	// Spawn an enemy
	if rl.IsKeyPressed(rl.KeySpace) {
		g.SpawnEnemy()
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

	// Draw a triangle towards the players attack direction
	// g.Player.DrawAttackTriangle(g.Player.CalculcateMeleeAttackArea(100, 300))

	g.DestroyProjectiles()
	g.DestroyEnemy()

}

func (g *Game) SpawnPlayer() {
	g.Player = NewPlayer("Dracula", 50, 100, 150, 100, 50)
}

func (g *Game) SpawnProjectile(x, y, radius, speed float32, direction rl.Vector2, color rl.Color, isHoming bool) {
	g.Projectiles = append(g.Projectiles, NewProjectile(x, y, radius, speed, direction, color, isHoming))
}

func (g *Game) DestroyProjectiles() {
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
