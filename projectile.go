package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Projectile struct {
	X         float32
	Y         float32
	Radius    float32
	Speed     float32
	Active    bool
	Direction rl.Vector2
}

type ProjectileSpawner interface {
	SpawnProjectile(x, y, radius, speed float32, direction rl.Vector2)
}

func NewProjectile(x, y, radius, speed float32, direction rl.Vector2) *Projectile {

	return &Projectile{
		X:         x,
		Y:         y,
		Radius:    radius,
		Speed:     speed,
		Active:    true,
		Direction: direction,
	}
}

func (p *Projectile) Update(g *Game) {

	if p.Active {
		p.X += p.Direction.X * p.Speed * rl.GetFrameTime()
		p.Y += p.Direction.Y * p.Speed * rl.GetFrameTime()

		// Check for collisions with enemies
		for _, enemy := range g.Enemies {
			if p.CollidesWith(enemy) {
				p.Active = false
				enemy.TakeDamage(g.Player.Damage)
				break
			}
		}

		p.Render()

		// Deactivate projectiles that are out of bounds
		if p.X < 0 || p.X > float32(rl.GetScreenWidth()) || p.Y < 0 || p.Y > float32(rl.GetScreenHeight()) {
			p.Active = false
		}
	}
}

func (p *Projectile) Render() {
	rl.DrawCircle(int32(p.X), int32(p.Y), p.Radius, rl.Black)
}

func (p *Projectile) CollidesWith(e *Enemy) bool {
	return rl.CheckCollisionCircleRec(
		rl.Vector2{
			X: p.X,
			Y: p.Y,
		},
		p.Radius,
		rl.Rectangle{
			X:      e.X,
			Y:      e.Y,
			Width:  float32(e.Width),
			Height: float32(e.Height),
		},
	)
}
