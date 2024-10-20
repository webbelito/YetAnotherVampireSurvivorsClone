package enemy

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SPAWN_DISTANCE = 100
)

type Enemy struct {
	Name   string
	X      float32
	Y      float32
	Width  int32
	Height int32
	SpawnX float32
	SpawnY float32
}

func NewEnemy(n string, w int32, h int32) *Enemy {

	e := &Enemy{
		Name:   n,
		X:      0,
		Y:      0,
		Width:  w,
		Height: h,
	}

	e.RandomizeSpawnPosition()

	return e
}

func (e *Enemy) Update() {
	e.Render()
}

func (e *Enemy) Render() {
	rl.DrawRectangle(int32(e.X), int32(e.Y), e.Width, e.Height, rl.Red)
}

func (e *Enemy) Spawn() {
	e.RandomizeSpawnPosition()
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
