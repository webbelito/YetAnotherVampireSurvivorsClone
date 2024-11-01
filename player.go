package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player interface {
	GetPosition() (float32, float32)
	GetName() string
}

type PlayerCharacter struct {
	Name                  string
	X                     float32
	Y                     float32
	PreviousPosition      rl.Vector2
	Width                 int32
	Height                int32
	Collider              rl.Rectangle
	ColliderAdjustment    rl.Vector2
	Speed                 float32
	directionX            int32
	directionY            int32
	targetDirection       rl.Vector2
	Texture               rl.Texture2D
	TextureSourceRect     rl.Rectangle
	TextureBasePos        rl.Vector2
	Health                float32
	MaxHealth             float32
	IsDead                bool
	BaseExperience        int32
	Experience            int32
	RequiredExperience    int32
	Level                 int32
	Damage                float32
	LastShotTime          float64
	ShootCooldown         float32
	LastMeleeTime         float64
	MeleeAttackBasePos    rl.Vector2
	MeleeAttackSourceRect rl.Rectangle
	MeleeCooldown         float32
	PlayMeleeAttack       bool
	LastHomingTime        float64
	HomingCooldown        float32
	Projectiles           *[]*Projectile
	AttackDirection       rl.Vector2
	PowerUps              []*PowerUp
	HUD                   *HUD

	// TODO: Implement an animator for playing animations
	// Idle Animation
	frameIndex   int
	frameTime    float32
	frameTimer   float32
	totalFrames  int
	framesWidth  int
	framesHeight int

	// Melee Animation
	FrameIndexMelee   int
	FrameTimeMelee    float32
	FrameTimerMelee   float32
	TotalFramesMelee  int
	FramesWidthMelee  int
	FramesHeightMelee int
}

func NewPlayer(n string, w int32, h int32, s float32, health float32, d float32) *PlayerCharacter {
	p := &PlayerCharacter{
		Name:                  n,
		X:                     float32(rl.GetScreenWidth()) / 2,
		Y:                     float32(rl.GetScreenHeight()) / 2,
		Width:                 w,
		Height:                h,
		Speed:                 s,
		directionX:            0,
		directionY:            0,
		targetDirection:       rl.Vector2{X: 1, Y: 0},
		Texture:               TextureAtlas,
		TextureSourceRect:     rl.NewRectangle(0, 64, 32, 64),
		Health:                health,
		MaxHealth:             health,
		BaseExperience:        10,
		Experience:            0,
		RequiredExperience:    CalculateXPForLevel(1, 10),
		Level:                 1,
		Damage:                d,
		LastShotTime:          0,
		ShootCooldown:         1,
		MeleeAttackSourceRect: rl.NewRectangle(0, 254, 32, 64),
		MeleeAttackBasePos:    rl.NewVector2(32, 0),
		LastMeleeTime:         0,
		MeleeCooldown:         2,
		LastHomingTime:        0,
		HomingCooldown:        5,
		PowerUps:              make([]*PowerUp, 0),
		frameIndex:            1,
		frameTime:             0.1,
		frameTimer:            0,

		// Melee Animation
		FrameIndexMelee: 1,
		FrameTimeMelee:  0.03,
		FrameTimerMelee: 0,
	}

	p.TextureBasePos = rl.NewVector2(0, 64)

	// Animation stuffs
	p.totalFrames = 8
	p.framesWidth = 32
	p.framesHeight = 64

	// Melee Animation
	p.TotalFramesMelee = 10
	p.FramesWidthMelee = 32
	p.FramesHeightMelee = 64

	p.TextureSourceRect = rl.NewRectangle(
		p.TextureBasePos.X,
		p.TextureBasePos.Y,
		float32(p.framesWidth),
		float32(p.framesHeight),
	)

	// Create the Collider
	p.ColliderAdjustment = rl.NewVector2(0, 9)
	p.Collider = rl.NewRectangle(p.X, p.Y, float32(p.Width-int32(p.ColliderAdjustment.X)), float32(p.Height-int32(p.ColliderAdjustment.Y)))

	return p
}

func (p *PlayerCharacter) Update(g *Game) {
	p.HandleInput()

	// TODO: Implement a better way to handle loadouts
	if p.Level >= 1 {
		p.Fire(g)
	}

	if p.Level >= 2 {
		p.Melee(g)
	}

	if p.Level >= 4 {
		p.ShootHoming(g)
	}

	// Check if we should Level Up
	if p.Experience >= p.RequiredExperience {
		p.LevelUp()
		g.ChangeGameState(LeveledUp)
	}
}

func (p *PlayerCharacter) FixedUpdate(g *Game) {
	p.HandleMovment()
	p.HandleColliders(g)
}

func (p *PlayerCharacter) HandleMovment() {

	// Store the last position for interpolation
	p.PreviousPosition = rl.Vector2{X: p.X, Y: p.Y}

	totalSpeed := p.Speed

	// Apply powerups
	for i := 0; i < len(p.PowerUps); i++ {

		if p.PowerUps[i].Type == Speed {
			// Example 50 += 0.5 * 50 = 75
			totalSpeed += p.PowerUps[i].SpeedIncreasePercentage * p.Speed
		}
	}

	p.X += float32(p.directionX) * totalSpeed * rl.GetFrameTime()
	p.Y += float32(p.directionY) * totalSpeed * rl.GetFrameTime()

}

func (p *PlayerCharacter) HandleInput() {

	// Handle horizontal movement
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.directionX = 1
		p.targetDirection = rl.Vector2{X: 1, Y: 0}
		p.TurnPlayerRight()
	} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.directionX = -1
		p.targetDirection = rl.Vector2{X: -1, Y: 0}
		p.TurnPlayerLeft()
	} else {
		p.directionX = 0
	}

	// Handle vertical movement
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.directionY = -1
		p.AttackDirection = rl.Vector2{X: 0, Y: -1}
	} else if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.directionY = 1
		p.AttackDirection = rl.Vector2{X: 0, Y: 1}
	} else {
		p.directionY = 0
	}
}

func (p *PlayerCharacter) HandleColliders(g *Game) {

	// Update the collider position
	p.Collider.X = p.X
	p.Collider.Y = p.Y

	// Check for collisions with walls
	if g.Level.CheckCollisions(p) {

		// Push the player back
		p.X -= float32(p.directionX) * p.Speed * rl.GetFrameTime()
		p.Y -= float32(p.directionY) * p.Speed * rl.GetFrameTime()
	}
}

func (p *PlayerCharacter) GetCollider() rl.Rectangle {
	return p.Collider
}
func (p *PlayerCharacter) TurnPlayerRight() {
	p.TextureSourceRect = rl.NewRectangle(0, 64, 32, 64)
}

func (p *PlayerCharacter) TurnPlayerLeft() {
	p.TextureSourceRect = rl.NewRectangle(0, 64, -32, 64)
}

func (p *PlayerCharacter) Attack(e Entity) {
	e.TakeDamage(p.Damage)
}

func (p *PlayerCharacter) DealDamage(e Entity) {

	totalDamage := p.Damage

	// Apply powerups
	for i := 0; i < len(p.PowerUps); i++ {

		if p.PowerUps[i].Type == Damage {
			// Example 50 += 0.5 * 50 = 75
			totalDamage += p.PowerUps[i].DamageIncreasePercentage * p.Damage
		}
	}

	e.TakeDamage(totalDamage)

}

func (p *PlayerCharacter) TakeDamage(damage float32) {

	rl.TraceLog(rl.LogDebug, "%s takes %f damage", p.Name, damage)

	if p.Health-damage <= 0 {
		p.Die()
		return
	}

	p.Health -= damage

}

func (p *PlayerCharacter) Die() {
	rl.TraceLog(rl.LogDebug, "Player %s has died", p.Name)

	p.IsDead = true
}

func (p *PlayerCharacter) Heal(amount float32) {

	// Check if the heal will exceed the max health
	if p.Health+amount > p.MaxHealth {

		rl.TraceLog(rl.LogDebug, "Healing player to max health")

		p.Health = p.MaxHealth
		return
	}

	rl.TraceLog(rl.LogDebug, "Healing player for %f", amount)

	p.Health += amount
}

func (p *PlayerCharacter) UpdateAnimation() {

	// Idle Animation
	if p.directionX == 0 && p.directionY == 0 {

		// Update the frame timer
		p.frameTimer += rl.GetFrameTime()

		if p.frameTimer >= p.frameTime {
			p.frameTimer = 0
			p.frameIndex++

			if p.frameIndex >= p.totalFrames {
				p.frameIndex = 1
			}

			p.TextureSourceRect.X = float32(p.frameIndex) * float32(p.framesWidth)
		}
	}

	// Melee Attack Animation
	if p.PlayMeleeAttack {

		// Update the frame timer
		p.FrameTimerMelee += rl.GetFrameTime()

		if p.FrameTimerMelee >= p.FrameTimeMelee {
			p.FrameTimerMelee = 0
			p.FrameIndexMelee++

			if p.FrameIndexMelee >= p.TotalFramesMelee {
				p.FrameIndexMelee = 1
				p.PlayMeleeAttack = false
			}

			p.MeleeAttackSourceRect.X = float32(p.FrameIndexMelee) * float32(p.FramesWidthMelee)

		}

	}

}

func (p *PlayerCharacter) Render(interpolation float64) {

	// Interpolate the player position
	interpolatedX := p.PreviousPosition.X*(1-float32(interpolation)) + p.X*float32(interpolation)
	interpolatedY := p.PreviousPosition.Y*(1-float32(interpolation)) + p.Y*float32(interpolation)

	// Draw the player collider
	// rl.DrawRectangle(int32(p.Collider.X), int32(p.Collider.Y), int32(p.Collider.Width), int32(p.Collider.Height), rl.Red)

	// Draw the player character

	rl.DrawTexturePro(
		p.Texture,
		p.TextureSourceRect,
		rl.NewRectangle(interpolatedX, interpolatedY, float32(p.Width), float32(p.Height)),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

	// Draw the melee attack
	if p.PlayMeleeAttack {

		// Set the melee attack animation direction
		if p.targetDirection.X > 0 {
			p.MeleeAttackSourceRect.Width = 32
			p.MeleeAttackBasePos.X = 32
		} else {
			p.MeleeAttackSourceRect.Width = -32
			p.MeleeAttackBasePos.X = -32
		}

		rl.DrawTexturePro(
			p.Texture,
			p.MeleeAttackSourceRect,
			rl.NewRectangle(p.X+p.MeleeAttackBasePos.X, p.Y+p.MeleeAttackBasePos.Y, float32(p.FramesWidthMelee), float32(p.FramesHeightMelee)),
			rl.Vector2{X: 16, Y: 0},
			0,
			rl.White,
		)

	}

}

func (p *PlayerCharacter) GetName() string {
	return p.Name
}

func (p *PlayerCharacter) GetPosition() (float32, float32) {
	return p.X, p.Y
}

// Shooting
func (p *PlayerCharacter) CanShoot() bool {
	return (p.LastShotTime >= float64(p.ShootCooldown))
}

func (p *PlayerCharacter) Fire(g *Game) {

	// Update the timer
	p.LastShotTime += float64(rl.GetFrameTime())

	if p.CanShoot() {

		// Reset the last shot time
		p.LastShotTime = 0

		// Calculate the target direction based on the player's movement
		direction := rl.Vector2Normalize(p.targetDirection)

		g.SpawnProjectile(float32(p.X+float32(p.Width)/2), float32(p.Y+float32(p.Height)/2), 5, 500, direction, rl.Black, false)

	}

}

// Melee attack
func (p *PlayerCharacter) CanMeleeAttack() bool {
	return (p.LastMeleeTime >= float64(p.MeleeCooldown))
}

func (p *PlayerCharacter) Melee(g *Game) {

	// Update the timer
	p.LastMeleeTime += float64(rl.GetFrameTime())

	if p.CanMeleeAttack() {

		// Reset the last melee attack time
		p.LastMeleeTime = 0

		// Set the melee attack flag to true
		p.PlayMeleeAttack = true

		// Calculate the melee attack area
		tip, baseLeftCorner, baseRightCorner := p.CalculcateMeleeAttackArea(100, 300)

		// Check for collisions with enemies

		for i := 0; i < len(g.Enemies); i++ {
			if rl.CheckCollisionPointTriangle(rl.Vector2{X: g.Enemies[i].X, Y: g.Enemies[i].Y}, tip, baseLeftCorner, baseRightCorner) {
				if !g.Enemies[i].IsDead {
					p.DealDamage(g.Enemies[i])
				}
			}

		}

		// Draw the attack triangle
		//p.DrawAttackTriangle(tip, baseLeftCorner, baseRightCorner)

	}

}

func (p *PlayerCharacter) DrawAttackTriangle(tip rl.Vector2, baseLeftCorner rl.Vector2, baseRightCorner rl.Vector2) {

	// Draw the triangle using the calculated points
	rl.DrawTriangle(
		tip,
		baseLeftCorner,
		baseRightCorner,
		rl.ColorAlpha(rl.Orange, 0.3),
	)
}

func (p *PlayerCharacter) CalculcateMeleeAttackArea(length float32, baseWidth float32) (rl.Vector2, rl.Vector2, rl.Vector2) {

	// Tip of the triangle (at player position)
	tipX := p.X + float32(p.Width)/2
	tipY := p.Y + float32(p.Height)/2

	// Normalize the attack direction
	normalizedDirection := rl.Vector2Normalize(p.targetDirection)

	// Normalize the attack towards the mouse
	//normalizedDirection := rl.Vector2Normalize(rl.Vector2{X: float32(rl.GetMouseX()) - tipX, Y: float32(rl.GetMouseY()) - tipY})

	// Calculate the base center (wide end of the triangle)
	baseCenterX := tipX + normalizedDirection.X*length
	baseCenterY := tipY + normalizedDirection.Y*length

	// Perpendicular vector to the attack direction
	perpX := -normalizedDirection.Y
	perpY := normalizedDirection.X

	// Calculate the two base corners of the triangle
	baseLeftCornerX := baseCenterX + perpX*baseWidth/2
	baseLeftCornerY := baseCenterY + perpY*baseWidth/2

	baseRightCornerX := baseCenterX - perpX*baseWidth/2
	baseRightCornerY := baseCenterY - perpY*baseWidth/2

	return rl.Vector2{X: tipX, Y: tipY}, rl.Vector2{X: baseLeftCornerX, Y: baseLeftCornerY}, rl.Vector2{X: baseRightCornerX, Y: baseRightCornerY}

}

// Homing attack
func (p *PlayerCharacter) CanHomingAttack() bool {
	return (p.LastHomingTime >= float64(p.HomingCooldown))
}

func (p *PlayerCharacter) ShootHoming(g *Game) {

	// Update the timer
	p.LastHomingTime += float64(rl.GetFrameTime())

	if p.CanHomingAttack() {

		// Reset the last homing attack time
		p.LastHomingTime = 0

		direction := rl.Vector2Normalize(rl.Vector2{X: p.X, Y: p.Y})

		g.SpawnProjectile(float32(p.X+float32(p.Width)/2), float32(p.Y+float32(p.Height)/2), 5, 500, direction, rl.Purple, true)

	}

}

func (p *PlayerCharacter) FindClosestEnemy(g *Game) *Enemy {

	// Set the initial distance to a high value
	closestDistance := float32(1000000)
	var closestEnemy *Enemy

	// Find the closest enemy
	for i := 0; i < len(g.Enemies); i++ {

		if g.Enemies[i].IsDead {
			continue
		}

		// Calculate the distance to the enemy
		distance := rl.Vector2Distance(rl.Vector2{X: p.X, Y: p.Y}, rl.Vector2{X: g.Enemies[i].X, Y: g.Enemies[i].Y})

		// Check if the enemy is closer
		if distance < closestDistance {
			closestDistance = distance
			closestEnemy = g.Enemies[i]
		}
	}

	return closestEnemy
}

// Powerups
func (p *PlayerCharacter) PowerUpUpdate(g *Game) {
	for i := 0; i < len(p.PowerUps); i++ {
		p.PowerUps[i].Update(g)
	}
}

// Experience
func (p *PlayerCharacter) GainExperience(amount int32) {

	// Add the experience
	p.Experience += amount

	// Check if the player leveled up
	if p.Experience >= p.RequiredExperience {
		p.LevelUp()
	}
}

func (p *PlayerCharacter) LevelUp() {

	baseXP := 10

	// Increase the level
	p.Level++

	// TODO: Implement level up bonuses
	/*
		// Increase the max health
		p.MaxHealth += 10

		// Increase the damage
		p.Damage += 5

		// Increase the speed
		p.Speed += 10
	*/

	// Reset the health
	p.Health = p.MaxHealth

	// Reset the experience
	p.Experience = 0

	// Increase the required experience by 50%
	p.RequiredExperience = CalculateXPForLevel(p.Level, int32(baseXP))

	// Show the level up message
	p.HUD.LeveledUp = true

	rl.TraceLog(rl.LogDebug, "Player %s has leveled up to level %d", p.Name, p.Level)
}

func CalculateXPForLevel(level int32, baseXP int32) int32 {
	return int32(float64(baseXP) * float64(level) * math.Pow(1.1, float64(level)))
}
