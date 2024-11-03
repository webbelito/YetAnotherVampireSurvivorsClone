package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SkillManager struct {
	AllSkills    map[string]*Skill
	ActiveSkills map[string]*Skill
}

func NewSkillManager() *SkillManager {
	return &SkillManager{
		AllSkills:    make(map[string]*Skill),
		ActiveSkills: make(map[string]*Skill),
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

			// If the skill is already active, try again
			// TODO: Make sure we don't get stuck in an infinite loop
			sm.SelectRandomSkill()
			return
		}
	}

	// Add the skill to the active skills list
	sm.SelectSkill(randomSkill)

}
