package player

import (
	"fmt"

	"github.com/webbelito/YetAnotherVampireSurvivorsClone/entity"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/projectile"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player interface {
	GetPosition() (float32, float32)
}

type PlayerCharacter struct {
	Name          string
	X             float32
	Y             float32
	Width         int32
	Height        int32
	Speed         float32
	directionX    int32
	directionY    int32
	Health        int32
	Damage        int32
	LastShotTime  float64
	ShootCooldown float32
	Projectiles   *[]*projectile.Projectile
}

func NewPlayer(n string, w int32, h int32, s float32, health int32, d int32) *PlayerCharacter {
	return &PlayerCharacter{
		Name:          n,
		X:             0,
		Y:             0,
		Width:         w,
		Height:        h,
		Speed:         s,
		directionX:    0,
		directionY:    0,
		Health:        health,
		Damage:        d,
		LastShotTime:  0,
		ShootCooldown: 2,
	}
}

func (p *PlayerCharacter) Update() {
	p.HandleInput()
	p.HandleMovment()
	p.Shoot()
	p.Render()
}

func (p *PlayerCharacter) HandleMovment() {
	p.X += float32(p.directionX) * p.Speed * rl.GetFrameTime()
	p.Y += float32(p.directionY) * p.Speed * rl.GetFrameTime()

	// Limit player movement to screen width
	if p.X <= 0 {
		p.X = 0
	} else if p.X >= float32(int32(rl.GetScreenWidth())-p.Width) {
		p.X = float32(int32(rl.GetScreenWidth()) - p.Width)
	}

	// Limit player movement to screen height
	if p.Y <= 0 {
		p.Y = 0
	}
	if p.Y >= float32(int32(rl.GetScreenHeight())-p.Height) {
		p.Y = float32(int32(rl.GetScreenHeight()) - p.Height)
	}

}

func (p *PlayerCharacter) HandleInput() {

	// Handle horizontal movement
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.directionX = 1
	} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.directionX = -1
	} else {
		p.directionX = 0
	}

	// Handle vertical movement
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.directionY = -1
	} else if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.directionY = 1
	} else {
		p.directionY = 0
	}
}

func (p *PlayerCharacter) Attack(e entity.Entity) {
	e.TakeDamage(p.Damage)
}

func (p *PlayerCharacter) TakeDamage(damage int32) {
	p.Health -= damage

	fmt.Println(p.Name, "took", damage, "damage. Remaining health:", p.Health)
}

func (p *PlayerCharacter) Render() {
	rl.DrawRectangle(int32(p.X), int32(p.Y), p.Width, p.Height, rl.Green)
}

func (p *PlayerCharacter) GetName() string {
	return p.Name
}

func (p *PlayerCharacter) GetPosition() (float32, float32) {
	return p.X, p.Y
}

// Shooting
func (p *PlayerCharacter) CanShoot() bool {
	currentTime := rl.GetTime()
	return (currentTime-p.LastShotTime >= float64(p.ShootCooldown))
}

func (p *PlayerCharacter) Shoot() {
	if p.CanShoot() {
		p.LastShotTime = rl.GetTime()
		mousePos := rl.GetMousePosition()
		direction := rl.Vector2{
			X: mousePos.X - float32(p.X),
			Y: mousePos.Y - float32(p.Y),
		}

		direction = rl.Vector2Normalize(direction)

		rl.TraceLog(rl.LogInfo, "Player shoot direction: %v\n", direction)

		// TODO: Spawn a new projectile
		//projectiles = append(projectiles, projectile.NewProjectile(float32(p.X), float32(p.Y), 5, 500, direction))

	}
}
