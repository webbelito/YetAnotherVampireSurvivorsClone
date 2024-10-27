package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	/*
		"github.com/webbelito/YetAnotherVampireSurvivorsClone/entity"
		"github.com/webbelito/YetAnotherVampireSurvivorsClone/player"
		"github.com/webbelito/YetAnotherVampireSurvivorsClone/projectile"
	*/)

const (
	SPAWN_DISTANCE = 100
)

type Enemy struct {
	Name   string
	X      float32
	Y      float32
	Width  int32
	Height int32
	Speed  float32
	SpawnX float32
	SpawnY float32
	Health float32
	Damage float32
	IsDead bool
}

func NewEnemy(n string, w int32, h int32, s float32, health float32, d float32) *Enemy {

	e := &Enemy{
		Name:   n,
		X:      0,
		Y:      0,
		Width:  w,
		Height: h,
		Speed:  s,
		Health: health,
		Damage: d,
		IsDead: false,
	}

	e.RandomizeSpawnPosition()

	return e

}

func (e *Enemy) Update(target Player) {
	playerPosX, playerPosY := target.GetPosition()

	e.MoveTowardsPlayer(playerPosX, playerPosY)
	e.Render()
}

func (e *Enemy) Render() {
	rl.DrawRectangle(int32(e.X), int32(e.Y), e.Width, e.Height, rl.Red)
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
