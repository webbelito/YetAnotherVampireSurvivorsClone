package main

type Skill struct {
	Name          string
	BaseDamage    float32
	BaseRange     float32
	BaseCooldown  float32
	CooldownTimer float32
	LastUsed      float32
	CurrentLevel  int
	MaxLevel      int
	IsProjectile  bool
	UpgradePath   []UpgradeEffect
}

type UpgradeEffect struct {
	AdditionalProjectiles int
	AdditionalDamage      float32
	DamageMultiplier      float32
	RangeMultiplier       float32
	CooldownReduction     float32
	IsPiercing            bool
}

func NewSkill(name string, damage float32, baseRange float32, cooldown float32, maxLevel int, isProjectile bool, upgradePath []UpgradeEffect) *Skill {
	return &Skill{
		Name:         name,
		BaseDamage:   damage,
		BaseRange:    baseRange,
		BaseCooldown: cooldown,
		CurrentLevel: 1,
		MaxLevel:     maxLevel,
		IsProjectile: isProjectile,
		UpgradePath:  upgradePath,
	}
}

func (s *Skill) Upgrade() {
	if s.CurrentLevel < s.MaxLevel {
		s.CurrentLevel++

		// TODO: Apply upgrade effects

	}
}

func (s *Skill) IsAvailable() bool {
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
