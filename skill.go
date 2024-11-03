package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO: Projectile stuffs, refactor?
type ProjectilePattern int

const (
	Single ProjectilePattern = iota
	Cone
	Surround
)

type Skill struct {
	Name              string
	BaseDamage        float32
	BaseRange         float32
	BaseCooldown      float32
	CooldownTimer     float32
	LastUsed          float32
	CurrentLevel      int
	MaxLevel          int
	ProjectilesCount  int
	ProjectileType    ProjectileType
	ProjectileRadius  float32
	ProjectileSpeed   float32
	ProjectilePattern ProjectilePattern
	UpgradePath       []UpgradeEffect
}
type UpgradeEffect struct {
	AdditionalProjectiles int
	AdditionalDamage      float32
	DamageMultiplier      float32
	RangeMultiplier       float32
	CooldownReduction     float32
	IsPiercing            bool
}

func NewSkill(name string, damage float32, baseRange float32, cooldown float32, maxLevel int, projectilesCount int, projectileType ProjectileType, projectileRadius float32, projectileSpeed float32, projectilePattern ProjectilePattern, upgradePath []UpgradeEffect) *Skill {
	return &Skill{
		Name:              name,
		BaseDamage:        damage,
		BaseRange:         baseRange,
		BaseCooldown:      cooldown,
		CurrentLevel:      1,
		MaxLevel:          maxLevel,
		ProjectilesCount:  projectilesCount,
		ProjectileType:    projectileType,
		ProjectileRadius:  projectileRadius,
		ProjectileSpeed:   projectileSpeed,
		ProjectilePattern: projectilePattern,
		UpgradePath:       upgradePath,
	}
}

func (s *Skill) Upgrade() {

	if !s.IsUpgradable() {
		return
	}

	if s.CurrentLevel < s.MaxLevel {
		s.CurrentLevel++

		// TODO: Apply upgrade effects

	}
}

func (s *Skill) IsUpgradable() bool {
	return s.CurrentLevel < s.MaxLevel
}

func (s *Skill) IsOnCooldown() bool {
	return s.CooldownTimer > 0
}

func (s *Skill) TriggerCooldown() {

	// Check if the skill is on cooldown
	if s.IsOnCooldown() {
		return
	}

	// Set the cooldown timer
	s.CooldownTimer = s.BaseCooldown
	s.LastUsed = 0
}

func (s *Skill) Use(g *Game) {

	// Check if the skill is on cooldown
	if s.IsOnCooldown() {
		return
	}

	// Trigger the cooldown
	s.TriggerCooldown()

	// Check if we have a projectile skill
	if s.ProjectilesCount > 0 {

		// Get the amount of projectiles to spawn
		amountOfProjectiles := s.UpgradePath[s.CurrentLevel-1].AdditionalProjectiles + s.ProjectilesCount

		if s.ProjectilePattern == Single {

			for i := 0; i < amountOfProjectiles; i++ {

				// Get the direction of the player
				direction := g.Player.targetDirection

				spawnPos := rl.Vector2{
					X: g.Player.X + float32(g.Player.Width)/2,
					Y: g.Player.Y + float32(g.Player.Height)/2,
				}

				// Ask the ProjectileManager to create a new projectile
				g.SpawnProjectile(s.ProjectileType, spawnPos.X, spawnPos.Y, s.BaseDamage, s.ProjectileSpeed, direction, rl.Red)
			}
		} else if s.ProjectilePattern == Cone {

			// Creates a cone of projectiles

			// Get the spread of the cone
			coneSpread := math.Pi / 4

			// Get the direction of the player
			targetDir := rl.Vector2Normalize(g.Player.targetDirection)

			// Calculate the base angle of the cone
			baseAngle := math.Atan2(float64(targetDir.Y), float64(targetDir.X))

			// Calculate the start and end angle of the cone
			startAngle := baseAngle - coneSpread/2
			endAngle := baseAngle + coneSpread/2

			var angleStep float64

			// Calculate the angle step
			if amountOfProjectiles > 1 {
				angleStep = (endAngle - startAngle) / float64(amountOfProjectiles-1)

				// Otherwise just use the base angle
			} else {
				angleStep = baseAngle
			}

			// Offset distance from the player's center
			offsetDistance := float32(20)

			// Get the spawn position of the projectiles
			playerCenter := rl.Vector2{
				X: g.Player.X + float32(g.Player.Width)/2,
				Y: g.Player.Y + float32(g.Player.Height)/2,
			}

			// For each projectile find a spawn point around the player
			for i := 0; i < amountOfProjectiles; i++ {

				// Calculate the angle of the projectile
				angle := startAngle + angleStep*float64(i)

				// Calculate the direction of the projectile
				direction := rl.Vector2{X: float32(math.Cos(angle)), Y: float32(math.Sin(angle))}

				// SpawnPos is the player's center plus the direction times the offset distance
				spawnPos := rl.Vector2{
					X: playerCenter.X + direction.X*offsetDistance,
					Y: playerCenter.Y + direction.Y*offsetDistance,
				}

				// Create a new projectile
				// Ask the ProjectileManager to create a new projectile
				g.SpawnProjectile(
					s.ProjectileType,
					spawnPos.X,
					spawnPos.Y,
					s.BaseDamage,
					s.ProjectileSpeed,
					direction,
					rl.Red,
				)

			}
		} else if s.ProjectilePattern == Surround {
			// Create a surround of projectiles around the player

			// Calculate the spread of 360 degrees
			fullCircle := math.Pi * 2

			// Calculate the angle step on the number of projectiles
			angleStep := fullCircle / float64(amountOfProjectiles)

			playerCenter := rl.Vector2{
				X: g.Player.X + float32(g.Player.Width)/2,
				Y: g.Player.Y + float32(g.Player.Height)/2,
			}

			// Offset distance from the player's center
			offsetDistance := float32(50)

			// For each projectile find a spawn point around the player
			for i := 0; i < amountOfProjectiles; i++ {

				// Calculate the angle of the projectile
				angle := angleStep * float64(i)

				// Calculate the direction of the projectile
				direction := rl.Vector2{X: float32(math.Cos(angle)), Y: float32(math.Sin(angle))}

				spawnPos := rl.Vector2{
					X: playerCenter.X + direction.X*offsetDistance,
					Y: playerCenter.Y + direction.Y*offsetDistance,
				}

				// Create a new projectile
				// Ask the ProjectileManager to create a new projectile
				g.SpawnProjectile(
					s.ProjectileType,
					spawnPos.X,
					spawnPos.Y,
					s.BaseDamage,
					s.ProjectileSpeed,
					direction,
					rl.Red,
				)

				rl.TraceLog(rl.LogInfo, "Spawned projectile at %v", spawnPos)

			}

			rl.TraceLog(rl.LogInfo, "Spawned %v projectiles", amountOfProjectiles)
		}
	}
}
