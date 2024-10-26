package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Golang enum for PowerUpType
// PowerUpType represents the type of powerup
type PowerUpType int

// PowerUpType constants
const (
	Heal PowerUpType = iota
	Speed
	Damage
)

// String returns the string representation of the PowerUpType
func (p PowerUpType) String() string {
	return [...]string{"Heal", "Speed", "Damage"}[p]
}

type PowerUp struct {
	Type              PowerUpType
	Position          rl.Vector2
	Active            bool
	PickedUp          bool
	Expired           bool
	TotalDuration     float32
	RemainingDuration float32
	Color             rl.Color
}

func (pu *PowerUp) RandomizePowerUpType() {

	// Do we need a sudo random number generator?
	powerUpTypeIndex := rl.GetRandomValue(0, 2)

	rl.TraceLog(rl.LogInfo, "PowerUpType Index: %d", powerUpTypeIndex)

	switch powerUpTypeIndex {
	case 0:
		pu.Type = Heal
		pu.Color = rl.Green
	case 1:
		pu.Type = Speed
		pu.Color = rl.Blue
	case 2:
		pu.Type = Damage
		pu.Color = rl.Red
	default:
		pu.Type = Heal
		pu.Color = rl.Green
	}
}

func NewPowerUp(d float32) *PowerUp {

	return &PowerUp{
		Active:        false,
		PickedUp:      false,
		Expired:       false,
		TotalDuration: d,
	}
}

func (pu *PowerUp) Update(g *Game) {

	// Tick down the remaining duration of the powerup
	if pu.Active && pu.PickedUp {
		pu.RemainingDuration -= rl.GetFrameTime()

		if pu.RemainingDuration <= 0 {
			pu.Expire(g.Player)
		}

	} else {

		if pu.CollidesWithPlayer(g.Player) {
			pu.PickUp(g)
		}

		pu.Render()
	}
}

func (pu *PowerUp) RandomizeSpawnPoint() rl.Vector2 {

	pos := rl.Vector2{
		X: float32(rl.GetRandomValue(0, int32(rl.GetScreenWidth()))),
		Y: float32(rl.GetRandomValue(0, int32(rl.GetScreenHeight()))),
	}

	return pos
}

func (pu *PowerUp) Render() {
	rl.DrawCircleV(pu.Position, 40, pu.Color)
}

func (pu *PowerUp) CollidesWithPlayer(p *PlayerCharacter) bool {
	return rl.CheckCollisionCircleRec(pu.Position, 40, rl.NewRectangle(p.X, p.Y, float32(p.Width), float32(p.Height)))
}

func (pu *PowerUp) PickUp(g *Game) {

	pu.PickedUp = true

	// Activate powerup
	pu.Activate(g)

	rl.TraceLog(rl.LogDebug, "PowerUp picked up by Player!")
}

func (pu *PowerUp) Activate(g *Game) {
	pu.Active = true
	pu.RemainingDuration = pu.TotalDuration

	// Add powerup to player
	g.Player.PowerUps = append(g.Player.PowerUps, pu)

}

func (pu *PowerUp) Expire(p *PlayerCharacter) {
	pu.Active = false
	pu.Expired = true
	pu.RemainingDuration = 0

	// Remove powerup from player
	for i := 0; i < len(p.PowerUps); i++ {
		if p.PowerUps[i] == pu {
			p.PowerUps = append(p.PowerUps[:i], p.PowerUps[i+1:]...)
			rl.TraceLog(rl.LogDebug, "PowerUp expired on Player!")
			break
		}
	}
}
