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

	g.Player.Update(g)

	for i := 0; i < len(g.Enemies); i++ {
		g.Enemies[i].Update(g.Player)
	}

	for i := 0; i < len(g.Projectiles); i++ {
		g.Projectiles[i].Update(g)
	}

	g.DestroyProjectiles()

}

func (g *Game) SpawnPlayer() {
	g.Player = NewPlayer("Dracula", 50, 100, 150, 100, 50)
}

func (g *Game) SpawnProjectile(x, y, radius, speed float32, direction rl.Vector2) {
	g.Projectiles = append(g.Projectiles, NewProjectile(x, y, radius, speed, direction))
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
	g.Enemies = append(g.Enemies, NewEnemy("Bat", 50, 50, 100, 100, 50))
}
