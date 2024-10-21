package projectile

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

var projectiles []Projectile

func SpawnProjectile(x, y, radius, speed float32, direction rl.Vector2) {

	projectile := Projectile{
		X:         x,
		Y:         y,
		Radius:    radius,
		Speed:     speed,
		Active:    true,
		Direction: direction,
	}

	projectiles = append(projectiles, projectile)
}

func Update() {

	for i := 0; i < len(projectiles); i++ {

		if projectiles[i].Active {
			projectiles[i].X += projectiles[i].Direction.X * projectiles[i].Speed * rl.GetFrameTime()
			projectiles[i].Y += projectiles[i].Direction.Y * projectiles[i].Speed * rl.GetFrameTime()
			projectiles[i].Render()
		}

		// Deactivate projectiles that are out of bounds
		if projectiles[i].X < 0 || projectiles[i].X > float32(rl.GetScreenWidth()) || projectiles[i].Y < 0 || projectiles[i].Y > float32(rl.GetScreenHeight()) {
			projectiles[i].Active = false
			// Remove the projectile from the slice
			projectiles = append(projectiles[:i], projectiles[i+1:]...)
		}
	}
}

func (p *Projectile) Render() {

	rl.DrawCircle(int32(p.X), int32(p.Y), p.Radius, rl.Black)
}
