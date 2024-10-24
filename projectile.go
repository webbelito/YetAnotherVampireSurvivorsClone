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
	Color     rl.Color
	IsHoming  bool
}

type ProjectileSpawner interface {
	SpawnProjectile(x, y, radius, speed float32, direction rl.Vector2)
}

func NewProjectile(x, y, radius, speed float32, direction rl.Vector2, color rl.Color, isHoming bool) *Projectile {

	return &Projectile{
		X:         x,
		Y:         y,
		Radius:    radius,
		Speed:     speed,
		Active:    true,
		Direction: direction,
		Color:     color,
		IsHoming:  isHoming,
	}
}

func (p *Projectile) Update(g *Game) {

	if p.Active {

		if p.IsHoming {
			enemy := p.FindClosestEnemy(g.Enemies)
			if enemy != nil {
				p.Direction = rl.Vector2Normalize(rl.Vector2Subtract(rl.Vector2{X: enemy.X, Y: enemy.Y}, rl.Vector2{X: p.X, Y: p.Y}))
			}
		}

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
	rl.DrawCircle(int32(p.X), int32(p.Y), p.Radius, p.Color)
}

func (p *Projectile) FindClosestEnemy(e []*Enemy) *Enemy {
	var cEnemy *Enemy
	var closestDistance float32 = 1000000

	for _, enemy := range e {

		if enemy.IsDead {
			break
		}

		distance := rl.Vector2Distance(rl.Vector2{X: p.X, Y: p.Y}, rl.Vector2{X: enemy.X, Y: enemy.Y})
		if distance < closestDistance {
			closestDistance = distance
			cEnemy = enemy
		}
	}

	if cEnemy == nil {
		return nil
	}

	return cEnemy
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
