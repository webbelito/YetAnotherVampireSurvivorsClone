package enemy

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/entity"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/player"
)

const (
	SPAWN_DISTANCE = 100
)

var enemies []Enemy

type Enemy struct {
	Name   string
	X      float32
	Y      float32
	Width  int32
	Height int32
	Speed  float32
	SpawnX float32
	SpawnY float32
	Health int32
	Damage int32
}

func NewEnemy(n string, w int32, h int32, s float32, health int32, d int32) {

	e := &Enemy{
		Name:   n,
		X:      0,
		Y:      0,
		Width:  w,
		Height: h,
		Speed:  s,
		Health: health,
		Damage: d,
	}

	e.RandomizeSpawnPosition()

	enemies = append(enemies, *e)
}

func Update(p *player.Player) {

	for i := 0; i < len(enemies); i++ {
		enemies[i].MoveTowardsPlayer(p)
		enemies[i].Render()
	}
}

func (e *Enemy) Render() {
	rl.DrawRectangle(int32(e.X), int32(e.Y), e.Width, e.Height, rl.Red)
}

func (e *Enemy) Spawn() {
	e.RandomizeSpawnPosition()
}

func (e *Enemy) MoveTowardsPlayer(p *player.Player) {
	if e.X < p.X {
		e.X += 1
	} else if e.X > p.X {
		e.X -= 1
	}

	if e.Y < p.Y {
		e.Y += 1
	} else if e.Y > p.Y {
		e.Y -= 1
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

func (e *Enemy) Attack(entity entity.Entity) {
	entity.TakeDamage(e.Damage)
}

func (e *Enemy) TakeDamage(d int32) {
	e.Health -= d

	fmt.Println(e.Name, "took", d, "damage. Remaining health:", e.Health)
}
