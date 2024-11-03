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
	IsProjectile      bool
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

func NewSkill(name string, damage float32, baseRange float32, cooldown float32, maxLevel int, isProjectile bool, projectileType ProjectileType, projectileRadius float32, projectileSpeed float32, projectilePattern ProjectilePattern, upgradePath []UpgradeEffect) *Skill {
	return &Skill{
		Name:              name,
		BaseDamage:        damage,
		BaseRange:         baseRange,
		BaseCooldown:      cooldown,
		CurrentLevel:      1,
		MaxLevel:          maxLevel,
		IsProjectile:      isProjectile,
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
	if s.IsProjectile {

		// Get the amount of projectiles to spawn
		amountOfProjectiles := s.UpgradePath[s.CurrentLevel-1].AdditionalProjectiles + 1

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
			// Create a cone of projectiles

			coneSpread := math.Pi / 4

			targetDir := rl.Vector2Normalize(g.Player.targetDirection)

			baseAngle := math.Atan2(float64(targetDir.Y), float64(targetDir.X))

			startAngle := baseAngle - coneSpread/2
			endAngle := baseAngle + coneSpread/2

			var angleStep float64

			if amountOfProjectiles > 1 {
				angleStep = (endAngle - startAngle) / float64(amountOfProjectiles-1)
			} else {
				angleStep = baseAngle
			}

			spawnPos := rl.Vector2{
				X: g.Player.X + float32(g.Player.Width)/2,
				Y: g.Player.Y + float32(g.Player.Height)/2,
			}

			for i := 0; i < amountOfProjectiles; i++ {

				// Calculate the angle of the projectile
				angle := startAngle + angleStep*float64(i)

				direction := rl.Vector2{X: float32(math.Cos(angle)), Y: float32(math.Sin(angle))}

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
		}

		if s.ProjectilePattern == Surround {
			// Create a surround of projectiles around the player
		}

		// For each projectile find a spawn point around the player
		// Calculate the direction of the projectile

		// Create a new projectile
		// Ask the ProjectileManager to create a new projectile

	}

	rl.TraceLog(rl.LogInfo, "Using skill: %s", s.Name)

}
