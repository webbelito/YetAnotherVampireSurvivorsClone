package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SPAWN_DISTANCE = 100
)

type Enemy struct {
	Name              string
	X                 float32
	Y                 float32
	PreviousPosition  rl.Vector2
	Radius            float32
	Width             int32
	Height            int32
	Collider          rl.Rectangle
	Texture           rl.Texture2D
	TextureSourceRect rl.Rectangle
	TextureBasePos    rl.Vector2
	Speed             float32
	SpawnX            float32
	SpawnY            float32
	Health            float32
	Damage            float32
	LastAttackTime    float32
	AttackCooldown    float32
	AttackRange       float32
	IsDead            bool
	frameIndex        int
	frameTime         float32
	frameTimer        float32
	totalFrames       int
	framesWidth       int
	framesHeight      int
}

func NewEnemy(t rl.Texture2D, n string, w int32, h int32, health float32) *Enemy {

	e := &Enemy{
		Name:           n,
		X:              0,
		Y:              0,
		Width:          w,
		Height:         h,
		Texture:        t,
		Health:         health,
		LastAttackTime: 1,
		AttackCooldown: 1,
		IsDead:         false,
		frameIndex:     1,
		frameTime:      0.1,
		frameTimer:     0,
	}

	e.RandomizeSpawnPosition()

	// TODO: Create a proper enemy type struct
	switch n {
	case "Bat":

		//  Texture
		e.TextureBasePos = rl.NewVector2(0, 0)

		e.TextureSourceRect = rl.NewRectangle(
			e.TextureBasePos.X,
			e.TextureSourceRect.Y,
			float32(e.framesWidth),
			float32(e.framesHeight),
		)

		e.Damage = 5
		e.Speed = 100
		e.AttackRange = 15

		e.Radius = 16

		// Animation
		e.totalFrames = 8
		e.framesWidth = 32
		e.framesHeight = 32
	case "Pumpkin":

		// Texture
		e.TextureBasePos = rl.NewVector2(0, 32)

		e.TextureSourceRect = rl.NewRectangle(
			e.TextureBasePos.X,
			e.TextureSourceRect.Y,
			float32(e.framesWidth),
			float32(e.framesHeight),
		)

		e.AttackRange = 20
		e.Damage = 25
		e.Speed = 20
		e.Health = 200

		// Animation
		e.totalFrames = 1
		e.framesWidth = 32
		e.framesHeight = 32
	default:
		rl.TraceLog(rl.LogError, "NewEnemy: Default Case, Unknown enemy type")
	}

	e.Collider = rl.NewRectangle(e.X, e.Y, float32(e.Width), float32(e.Height))

	return e

}

func (e *Enemy) Update(p *PlayerCharacter) {
	e.UpdateAnimation()
}

func (e *Enemy) FixedUpdate(g *Game) {
	e.MoveTowardsPlayer(g.Player)
	e.HandleColliders()
}

func (e *Enemy) HandleColliders() {

	// Update the collider position
	e.Collider.X = e.X
	e.Collider.Y = e.Y
}

func (e *Enemy) GetCollider() rl.Rectangle {
	return e.Collider
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

func (e *Enemy) Render(interpolation float64) {

	// Interpolate the enemy position
	interpolatedX := e.PreviousPosition.X*(1-float32(interpolation)) + e.X*float32(interpolation)
	interpolatedY := e.PreviousPosition.Y*(1-float32(interpolation)) + e.Y*float32(interpolation)

	e.TextureSourceRect = rl.NewRectangle(
		float32(e.frameIndex*e.framesWidth),
		float32(e.TextureBasePos.Y),
		float32(e.framesWidth),
		float32(e.framesHeight),
	)

	rl.DrawTexturePro(
		e.Texture,
		e.TextureSourceRect,
		rl.NewRectangle(interpolatedX, interpolatedY, float32(e.Width), float32(e.Height)),
		rl.NewVector2(16, 16),
		0,
		rl.White,
	)
}

func (e *Enemy) Spawn() {
	e.RandomizeSpawnPosition()
}

// TODO: Implement a better way of handling enemy collisions
// This function is O(n^2) and should be optimized
// Look at the QuadTree data structure or other spatial partitioning algorithms
func ResolveEnemyCollisions(g *Game) {

	for i := 0; i < len(g.Enemies); i++ {
		for j := i + 1; j < len(g.Enemies); j++ {
			e1 := g.Enemies[i]
			e2 := g.Enemies[j]

			// Calculate the squared distance to avoid using the square root
			distanceX := e1.X - e2.X
			distanceY := e1.Y - e2.Y
			distanceSquared := distanceX*distanceX + distanceY*distanceY
			minDistance := e1.Radius + e2.Radius
			minDistanceSquared := minDistance * minDistance

			// Only proceed if the enemies are colliding
			if distanceSquared < minDistanceSquared {

				// Compute distance and overlap
				distance := float32(math.Sqrt(float64(distanceSquared)))
				overlap := (minDistance - distance) / 2

				// Normalize the separation vector to get separation
				separation := rl.Vector2{
					X: overlap * distanceX / distance,
					Y: overlap * distanceY / distance,
				}

				// Move the enemy away from the collision point
				e1.X += separation.X
				e1.Y += separation.Y
				e2.X -= separation.X
				e2.Y -= separation.Y

			}
		}
	}
}

func (e *Enemy) MoveTowardsPlayer(p *PlayerCharacter) {

	e.PreviousPosition = rl.NewVector2(e.X, e.Y)

	posX, posY := p.GetPosition()

	// Check if we're in range of the player to attack
	if e.IsPlayerInAttackRange(p) {
		e.Attack(p)

		// Calculate the distance between the player and the enemy
		distance := rl.Vector2Distance(rl.NewVector2(e.X, e.Y), rl.NewVector2(p.X, p.Y))

		// If we're close enough to the player, don't move
		// TODO: Replace the hardcoded value with a variable
		if distance < e.AttackRange-10 {
			return
		}
	}

	// Move towards the player on the X and Y axis
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

func (e *Enemy) IsPlayerInAttackRange(p *PlayerCharacter) bool {

	return rl.CheckCollisionCircleRec(
		rl.Vector2{X: e.X, Y: e.Y},
		e.AttackRange,
		rl.NewRectangle(p.X, p.Y, 32, 32),
	)
}

func (e *Enemy) CanMeleeAttack() bool {
	return (e.LastAttackTime >= e.AttackCooldown)
}

func (e *Enemy) Attack(entity Entity) {

	// Update the last attack time
	e.LastAttackTime += rl.GetFrameTime()

	if e.CanMeleeAttack() {

		// Check if the entity is a player
		if player, ok := entity.(*PlayerCharacter); ok {
			player.TakeDamage(e.Damage)

			rl.TraceLog(rl.LogDebug, "Enemy %s is attacking %s", e.Name, entity.GetName())

		}

		// Reset the attack timer
		e.LastAttackTime = 0
	}

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
