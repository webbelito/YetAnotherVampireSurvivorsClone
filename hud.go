package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type HUD struct {
	Player            *PlayerCharacter
	LeveledUp         bool
	LeveledUpDuration float32
}

func NewHUD(p *PlayerCharacter) *HUD {
	return &HUD{
		Player:            p,
		LeveledUp:         false,
		LeveledUpDuration: 2,
	}
}

// TODO: Check if we really need the game object
func (h *HUD) Render(g *Game) {
	h.RenderHealthBar()
	h.RenderExperienceBar()
	h.RenderGameTime(g)

	if h.LeveledUp && h.LeveledUpDuration > 0 {
		h.RenderLeveledUp()
	}
}

func (h *HUD) RenderHealthBar() {

	maxHealth := h.Player.MaxHealth
	currentHealth := h.Player.Health

	barWidth := int32(200)
	barHeight := int32(20)

	barYOffset := int32(150)

	posX := int32(rl.GetScreenWidth()/2 - int(barWidth)/2 + int(h.Player.Width/2))
	posY := int32(rl.GetScreenHeight()/2 + int(barHeight)/2 + int(barYOffset))

	// Create the background of the health bar
	rl.DrawRectangle(posX, posY, barWidth, barHeight, rl.Black)

	healthPercentage := (currentHealth / maxHealth)
	healthBarWidth := int32(float32(barWidth) * healthPercentage)

	// Create the health bar
	rl.DrawRectangle(posX, posY, healthBarWidth, barHeight, rl.Red)

	// Draw the health text
	healthText := fmt.Sprintf("HP: %.0f/%.0f", currentHealth, maxHealth)

	textXPosition := int32(10)
	textYPosition := int32(60)

	rl.DrawText(healthText, textXPosition, textYPosition, 20, rl.Red)
}

func (h *HUD) RenderExperienceBar() {

	currentExperience := h.Player.Experience
	maxExperience := h.Player.RequiredExperience
	level := h.Player.Level

	posX := int32(0)
	posY := int32(0)

	barWidth := int32(rl.GetScreenWidth())
	barHeight := int32(50)

	// Create the background of the experience bar
	rl.DrawRectangle(posX, posY, barWidth, barHeight, rl.ColorAlpha(rl.Black, 0.5))

	experiencePercentage := (float32(currentExperience) / float32(maxExperience))
	experienceBarWidth := int32(float32(barWidth) * experiencePercentage)

	// Create the experience bar
	rl.DrawRectangle(posX, posY, experienceBarWidth, barHeight, rl.Blue)

	// Draw the experience text
	experienceText := fmt.Sprintf("Level: %d Experience: %d/%d", level, currentExperience, maxExperience)

	textXPosition := int32(rl.GetScreenWidth() - 300)
	textYPosition := int32(barHeight + 10)

	rl.DrawText(experienceText, textXPosition, textYPosition, 20, rl.Black)

}

func (h *HUD) RenderLeveledUp() {

	// Decrease the duration of the level up text
	if h.LeveledUpDuration >= 0 {
		h.LeveledUpDuration -= rl.GetFrameTime()
	}

	// Draw the level up text
	levelUpText := fmt.Sprintf("Level Up! %d", h.Player.Level)

	textXPosition := int32(rl.GetScreenWidth()/2 - 200)
	textYPosition := int32(rl.GetScreenHeight()/2 - 300)

	rl.DrawText(levelUpText, textXPosition, textYPosition, 64, rl.Yellow)

	if h.LeveledUpDuration <= 0 {

		// Reset the level up text
		h.LeveledUp = false
		h.LeveledUpDuration = 2

	}
}

func (h *HUD) RenderGameTime(g *Game) {

	gameTime := ""

	displaySeconds := false
	displayMinutes := false
	displayHours := false

	// Format the game time into hours, minutes and seconds
	gameTimeHours := int(g.GameTime / 3600)
	gameTimeMinutes := int((g.GameTime - float32(gameTimeHours*3600)) / 60)
	gameTimeSeconds := int(g.GameTime - float32(gameTimeHours)*3600 - float32(gameTimeMinutes)*60)

	// Show only seconds if the game time is less than 1 minute
	if gameTimeHours == 0 && gameTimeMinutes == 0 {
		displaySeconds = true
		gameTime = fmt.Sprintf("%02d", gameTimeSeconds)
	} else if gameTimeHours == 0 {
		displayMinutes = true
		gameTime = fmt.Sprintf("%02d:%02d", gameTimeMinutes, gameTimeSeconds)
	} else {
		displayHours = true
		gameTime = fmt.Sprintf("%02d:%02d:%02d", gameTimeHours, gameTimeMinutes, gameTimeSeconds)
	}

	textXPosition := int32(0)
	textYPosition := int32(55)

	if displaySeconds && !displayMinutes && !displayHours {

		textXPosition = int32(rl.GetScreenWidth() / 2)

	} else if displaySeconds && displayMinutes && !displayHours {

		textXPosition = int32(rl.GetScreenWidth()/2 - 25)

	} else {
		textXPosition = int32(rl.GetScreenWidth()/2 - 50)
	}

	rl.DrawText(gameTime, textXPosition, textYPosition, 40, rl.White)
}
