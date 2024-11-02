package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ExperienceGem struct {
	Position          rl.Vector2
	Amount            int32
	Active            bool
	Texture           rl.Texture2D
	Rec               rl.Rectangle
	TextureSourceRect rl.Rectangle
}

func NewExperienceGem(x float32, y float32, a int32) *ExperienceGem {
	return &ExperienceGem{
		Position:          rl.NewVector2(x, y),
		Amount:            a,
		Active:            true,
		Texture:           TextureAtlas,
		Rec:               rl.NewRectangle(x, y, 32, 32),
		TextureSourceRect: rl.NewRectangle(0, 384, 32, 32),
	}
}

func (eg *ExperienceGem) FixedUpdate(g *Game) {

	// Update the rectangle position
	// Do we need to do this more than once?
	if eg.Active {
		if eg.Position.X != eg.Rec.X || eg.Position.Y != eg.Rec.Y {
			eg.Rec.X = eg.Position.X
			eg.Rec.Y = eg.Position.Y
		}
	}

	// Move the experience gem towards the player
	eg.MoveTowardsPlayer(g.Player)

	// Check for collisions with the player
	if eg.CheckCollision(g.Player) {
		eg.RewardExperience(g.Player)
	}

}

func (eg *ExperienceGem) CheckCollision(p *PlayerCharacter) bool {
	if eg.Active {

		if rl.CheckCollisionRecs(p.Collider, eg.Rec) {
			return true
		}
	}

	return false
}

func (eg *ExperienceGem) MoveTowardsPlayer(p *PlayerCharacter) {
	if eg.Active && rl.CheckCollisionCircleRec(rl.Vector2{X: p.X, Y: p.Y}, float32(p.ExperienceGemCollectRadius), eg.Rec) {
		eg.Position = rl.Vector2Lerp(eg.Position, rl.Vector2{X: p.X, Y: p.Y}, 0.05)
	}
}

func (eg *ExperienceGem) RewardExperience(p *PlayerCharacter) {
	if eg.Active {
		p.Experience += int32(eg.Amount)
		eg.Active = false
	}
}

func (eg *ExperienceGem) Render() {
	if eg.Active {

		//rl.DrawRectangle(int32(eg.Position.X), int32(eg.Position.Y), 32, 32, rl.Red)

		rl.DrawTexturePro(
			TextureAtlas,
			eg.TextureSourceRect,
			rl.NewRectangle(eg.Position.X, eg.Position.Y, 32, 32),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)
	}
}
