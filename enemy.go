package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SPAWN_DISTANCE = 100
)

type Enemy struct {
	Name              string
	X                 float32
	Y                 float32
	Width             int32
	Height            int32
	Texture           rl.Texture2D
	TextureSourceRect rl.Rectangle
	Speed             float32
	SpawnX            float32
	SpawnY            float32
	Health            float32
	Damage            float32
	IsDead            bool
	frameIndex        int
	frameTime         float32
	frameTimer        float32
	totalFrames       int
	framesWidth       int
	framesHeight      int
}

func NewEnemy(t rl.Texture2D, n string, w int32, h int32, health float32, d float32) *Enemy {

	e := &Enemy{
		Name:         n,
		X:            0,
		Y:            0,
		Width:        w,
		Height:       h,
		Texture:      t,
		Health:       health,
		Damage:       d,
		IsDead:       false,
		frameIndex:   0,
		frameTime:    0.1,
		frameTimer:   0,
		totalFrames:  8,
		framesWidth:  32,
		framesHeight: 32,
	}

	e.RandomizeSpawnPosition()

	// TODO: Create a proper enemy type struct
	switch n {
	case "Bat":
		e.TextureSourceRect = rl.NewRectangle(0, 0, 32, 32)
		e.Speed = 100
	case "Pumpkin":
		e.TextureSourceRect = rl.NewRectangle(0, 32, 32, 32)
		e.Speed = 20
		e.Health = 200
	default:
		rl.TraceLog(rl.LogError, "NewEnemy: Default Case, Unknown enemy type")
	}

	return e

}

func (e *Enemy) Update(target Player) {
	playerPosX, playerPosY := target.GetPosition()

	e.MoveTowardsPlayer(playerPosX, playerPosY)

	e.UpdateAnimation()

	e.Render()
}

func (e *Enemy) UpdateAnimation() {

	e.frameTimer += rl.GetFrameTime()

	if e.frameTimer >= e.frameTime {
		e.frameTimer = 0
		e.frameIndex++

		// Loop the animation
		if e.frameIndex >= e.totalFrames {
			e.frameIndex = 0
		}
	}

}

func (e *Enemy) Render() {

	e.TextureSourceRect = rl.NewRectangle(
		float32(e.frameIndex*e.framesWidth),
		0,
		float32(e.framesWidth),
		float32(e.framesHeight),
	)

	rl.DrawTexturePro(
		e.Texture,
		e.TextureSourceRect,
		rl.NewRectangle(e.X, e.Y, float32(e.Width), float32(e.Height)),
		rl.NewVector2(16, 16),
		0,
		rl.White,
	)
}

func (e *Enemy) Spawn() {
	e.RandomizeSpawnPosition()
}

func (e *Enemy) MoveTowardsPlayer(posX, posY float32) {

	if e.X < posX {
		e.X += 1 * e.Speed * rl.GetFrameTime()
	} else if e.X > posX {
		e.X -= 1 * e.Speed * rl.GetFrameTime()
	}

	if e.Y < posY {
		e.Y += 1 * e.Speed * rl.GetFrameTime()
	} else if e.Y > posY {
		e.Y -= 1 * e.Speed * rl.GetFrameTime()
	}
}

func (e *Enemy) RandomizeSpawnPosition() {

	// Randomize if we'll spawn left or right
	if rl.GetRandomValue(0, 1000) <= 500 {
		e.X = float32(rl.GetRandomValue(0, SPAWN_DISTANCE))
	} else {
		e.X = float32(rl.GetRandomValue(int32(rl.GetScreenWidth())-e.Width-SPAWN_DISTANCE, int32(rl.GetScreenWidth()-int(e.Width))))
	}

	// Randomize if we'll spawn up or down
	if rl.GetRandomValue(0, 1000) <= 500 {
		e.Y = float32(rl.GetRandomValue(0, SPAWN_DISTANCE))
	} else {
		e.Y = float32(rl.GetRandomValue(int32(rl.GetScreenHeight())-e.Height-SPAWN_DISTANCE, int32(rl.GetScreenHeight()-int(e.Height))))
	}

}

func (e *Enemy) GetPosition() (float32, float32) {
	return e.X, e.Y
}

func (e *Enemy) GetName() string {
	return e.Name
}

func (e *Enemy) Attack(entity Entity) {
	entity.TakeDamage(e.Damage)
}

func (e *Enemy) TakeDamage(d float32) {

	rl.TraceLog(rl.LogDebug, "%s takes %f damage", e.GetName(), d)

	if e.Health-d <= 0 {
		e.Die()
		return
	}

	e.Health -= d
}

func (e *Enemy) Die() {
	e.IsDead = true

	rl.TraceLog(rl.LogDebug, "Enemy %s has died", e.Name)
}

func (e *Enemy) Heal(amount float32) {
	e.Health += amount
}

func CheckCollisionAABB(p Projectile, e *Enemy) bool {

	// Check if the projectile is inside the enemy with radius of the projectile
	return p.X < e.X+float32(e.Width) && p.X > e.X && p.Y > e.Y && p.Y < e.Y+float32(e.Height)

}
