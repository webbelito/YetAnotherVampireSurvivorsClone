package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SkillManager struct {
	AllSkills       map[string]*Skill
	ActiveSkills    map[string]*Skill
	MaxActiveSkills int
}

func NewSkillManager() *SkillManager {
	return &SkillManager{
		AllSkills:       make(map[string]*Skill),
		ActiveSkills:    make(map[string]*Skill),
		MaxActiveSkills: 8,
	}
}

func (sm *SkillManager) AddSkill(skill *Skill) {
	sm.AllSkills[skill.Name] = skill
}

func (sm *SkillManager) SelectSkill(skill *Skill) {

	if !sm.SkillExists(skill.Name) {
		rl.TraceLog(rl.LogError, "Trying to add a skill that doesn't exist to the active skills list")
		return
	}

	if sm.HasSkill(skill.Name) {
		rl.TraceLog(rl.LogError, "Trying to add a skill that is already active")
		return
	}

	if sm.IsAtMaxActiveSkills() {
		rl.TraceLog(rl.LogError, "Trying to add a skill when the max active skills limit has been reached")
		return
	}

	// Add the skill to the active skills list
	sm.ActiveSkills[skill.Name] = skill
}

func (sm *SkillManager) SkillExists(skillName string) bool {
	_, exists := sm.AllSkills[skillName]
	return exists
}

func (sm *SkillManager) HasSkill(skillName string) bool {
	_, exists := sm.ActiveSkills[skillName]
	return exists
}

func (sm *SkillManager) GetSkill(skillName string) *Skill {
	return sm.AllSkills[skillName]
}

func (sm *SkillManager) GetActiveSkill(skillName string) *Skill {
	return sm.ActiveSkills[skillName]
}

func (sm *SkillManager) IsAtMaxActiveSkills() bool {
	return sm.GetActiveSkillsCount() >= sm.MaxActiveSkills
}

func (sm *SkillManager) GetSkillByIndex(index int32) *Skill {
	i := int32(0)
	for _, skill := range sm.AllSkills {
		if i == index {
			return skill
		}
		i++
	}

	return nil
}

func (sm *SkillManager) GetActiveSkillByIndex(index int32) *Skill {
	i := int32(0)
	for _, skill := range sm.ActiveSkills {
		if i == index {
			return skill
		}
		i++
	}

	return nil
}

func (sm *SkillManager) GetActiveSkillsCount() int {
	return len(sm.ActiveSkills)
}

func (sm *SkillManager) SelectRandomSkill() {

	// Create a random number between 1 and the number of skills
	randomSkillNumber := rl.GetRandomValue(1, int32(len(sm.AllSkills)))

	randomSkillIndex := randomSkillNumber - 1

	// Get the skill at the random index
	randomSkill := sm.GetSkillByIndex(randomSkillIndex)

	if sm.HasSkill(randomSkill.Name) {

		// Check if we can upgrade the skill
		if randomSkill.IsUpgradable() {
			randomSkill.Upgrade()
		} else {

			// TODO: Make sure we don't get stuck in an infinite loop
			// Check if all skills are active
			if sm.GetActiveSkillsCount() == len(sm.AllSkills) {
				rl.TraceLog(rl.LogInfo, "All skills are active")
				return
			}

			// Try again
			sm.SelectRandomSkill()
			return
		}
	}

	// Add the skill to the active skills list
	sm.SelectSkill(randomSkill)

}

func (sm *SkillManager) Update(g *Game) {

	// Update the cooldowns of the active skills
	for _, skill := range sm.ActiveSkills {
		if skill.IsOnCooldown() {
			skill.CooldownTimer -= rl.GetFrameTime()
		}
	}
}
