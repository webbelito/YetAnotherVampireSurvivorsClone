package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const FixedUpdateRate float32 = 60.0

type Game struct {
	currentGameState GameState
	Player           *PlayerCharacter
	Level            *Level
	Camera           rl.Camera2D
	Enemies          []*Enemy
	Projectiles      []*Projectile
	PowerUps         []*PowerUp
	ExperienceGems   []*ExperienceGem
	FixedDeltaTime   float32
	LastFixedUpdate  time.Time
	LastFrameTime    time.Time
	GameTime         float32
	WaveManager      *WaveManager
	SkillManager     *SkillManager
	IsPaused         bool
	IsPlayingMusic   bool
}

func NewGame() *Game {

	grid := []string{
		"@^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^@",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"[..........................................................]",
		"[........X.........................................X.......]",
		"[..........................................................]",
		"[..........................................................]",
		"@__________________________________________________________@",
	}

	level := NewLevel(grid)
	camera := rl.Camera2D{
		Offset: rl.Vector2{
			X: float32(rl.GetScreenWidth())/2 - 16,
			Y: float32(rl.GetScreenHeight()) / 2,
		},
		Target: rl.Vector2{
			X: float32(rl.GetScreenWidth())/2 - 16,
			Y: float32(rl.GetScreenHeight()) / 2,
		},
		Rotation: 0,
		Zoom:     2,
	}

	// Wave Manager stuffs
	waveManager := &WaveManager{
		Waves: []Wave{
			// Create 20 waves of enemies
			{EnemyCounts: map[EnemyType]int{Bat: 50}, SpawnInterval: 0.5, Duration: 10.0},
			{EnemyCounts: map[EnemyType]int{Bat: 50}, SpawnInterval: 0.3, Duration: 20.0},
			{EnemyCounts: map[EnemyType]int{Bat: 50}, SpawnInterval: 0.1, Duration: 30.0},
			{EnemyCounts: map[EnemyType]int{Bat: 50, Pumpkin: 50}, SpawnInterval: 0.1, Duration: 40.0},
			{EnemyCounts: map[EnemyType]int{Bat: 100, Pumpkin: 100}, SpawnInterval: 0.1, Duration: 50.0},
		},
	}

	// TODO: Load the skills
	skillManager := NewSkillManager()

	// Play the background music
	rl.PlayMusicStream(SoundTrack)

	return &Game{
		currentGameState: MainMenu,
		Player:           nil,
		Level:            level,
		Camera:           camera,
		Enemies:          make([]*Enemy, 0),
		Projectiles:      make([]*Projectile, 0),
		PowerUps:         make([]*PowerUp, 0),
		FixedDeltaTime:   1.0 / FixedUpdateRate,
		LastFixedUpdate:  time.Now(),
		WaveManager:      waveManager,
		SkillManager:     skillManager,
		IsPaused:         true,
		IsPlayingMusic:   true,
	}

}

// TODO: Remove this with a proper way to handle frame updates
var fixedFrameCounter int

func (g *Game) Update() {

	if g.Player != nil && g.Player.IsDead {
		g.ChangeGameState(GameOver)
	}

	// TODO: Create a better way of spawning the player
	if g.Player == nil {
		g.SpawnPlayer()
	}

	// Initialize the HUD
	if g.Player.HUD == nil {
		g.Player.HUD = NewHUD(g.Player)
	}

	// Check if we need to create all the skills
	if g.SkillManager != nil && len(g.SkillManager.AllSkills) == 0 {
		g.CreateAllSkills()
	}

	// Spawn an enemy
	if rl.IsKeyPressed(rl.KeyB) {

		for i := 0; i < 24; i++ {
			g.SpawnBat()
		}

		g.SpawnBat()
	}

	// Spawn a pumpkin
	if rl.IsKeyPressed(rl.KeySpace) {
		g.SpawnPumpkin()
	}

	// Spawn a powerup
	if rl.IsKeyPressed(rl.KeyP) {
		g.SpawnPowerUp()
	}

	// Player takes damage
	if rl.IsKeyPressed(rl.KeyK) {
		g.Player.TakeDamage(10)
	}

	// Player heals
	if rl.IsKeyPressed(rl.KeyH) {
		g.Player.Heal(10)
	}

	// Player gains experience
	if rl.IsKeyPressed(rl.KeyE) {
		g.Player.GainExperience(10)
	}

	if rl.IsKeyPressed(rl.KeyM) {
		if g.IsPlayingMusic {
			rl.PauseMusicStream(SoundTrack)
		} else {
			rl.ResumeMusicStream(SoundTrack)
		}

		g.IsPlayingMusic = !g.IsPlayingMusic
	}

	// Update player
	g.Player.Update(g)

	// Update enemies
	for i := 0; i < len(g.Enemies); i++ {

		if g.Enemies[i].IsDead {
			continue
		}

		g.Enemies[i].Update(g.Player)
	}

	// Update projectiles
	for i := 0; i < len(g.Projectiles); i++ {

		if !g.Projectiles[i].Active {
			continue
		}

		g.Projectiles[i].Update(g)
	}

	// Update powerups
	for i := 0; i < len(g.PowerUps); i++ {
		g.PowerUps[i].Update(g)
	}

	// Update Active Skills cooldown
	g.SkillManager.Update(g)

	// Play the next frame of the music
	rl.UpdateMusicStream(SoundTrack)

	g.DestroyProjectiles()
	g.DestroyEnemy()
	g.DestroyPowerUp()

	// Update the camera
	g.Camera.Target = rl.Vector2{
		X: g.Player.X,
		Y: g.Player.Y,
	}

}

func (g *Game) FixedUpdate() {

	fixedFrameCounter++

	// Player
	g.Player.FixedUpdate(g)

	// Resolve collisions every 3rd frame
	if fixedFrameCounter%3 == 0 {
		ResolveEnemyCollisions(g)
	}

	// Enemies
	for i := 0; i < len(g.Enemies); i++ {
		g.Enemies[i].FixedUpdate(g)
	}

	// Experience Gems
	for i := 0; i < len(g.ExperienceGems); i++ {
		g.ExperienceGems[i].FixedUpdate(g)
	}

	// Wave Manager
	if g.WaveManager.CurrentWaveIndex < len(g.WaveManager.Waves) {
		waveManager := g.WaveManager
		currentWave := waveManager.Waves[waveManager.CurrentWaveIndex]

		waveManager.TimeSinceLastSpawn += g.FixedDeltaTime
		waveManager.TimeSinceLastWave += g.FixedDeltaTime

		if waveManager.TimeSinceLastWave < currentWave.Duration {
			waveManager.SpawnEnemies(g)
		} else {
			waveManager.TimeSinceLastWave = 0
			waveManager.TimeSinceLastSpawn = 0
			waveManager.CurrentWaveIndex++
		}
	}

	// Check win condidition
	if g.WaveManager.CurrentWaveIndex >= len(g.WaveManager.Waves) && len(g.Enemies) == 0 {
		g.ChangeGameState(Victory)
	}
}

func (g *Game) Render(interpolation float64) {
	// Level stuffs here
	g.Level.Render()

	// Player stuffs here
	g.Player.UpdateAnimation()
	g.Player.Render(interpolation)

	// Enemy stuffs here
	for _, e := range g.Enemies {
		e.UpdateAnimation()
		e.Render(interpolation)
	}

	// Projectile stuffs here
	for _, p := range g.Projectiles {
		p.Render(interpolation)
	}

	// Experience Gem stuffs here
	for _, eg := range g.ExperienceGems {
		eg.Render()
	}
}

func (g *Game) Run() {

	// Interpolation
	var accumulatedTime float64 = 0.0

	for !rl.WindowShouldClose() {

		now := time.Now()
		deltaTime := time.Since(g.LastFrameTime).Seconds()
		g.LastFrameTime = now

		// Cap the deltaTime to 100ms
		if deltaTime > 0.1 {
			deltaTime = 0.1
		}

		//g.LastFixedUpdate = now

		// Do not update the game if it is paused
		if g.currentGameState != Paused {
			accumulatedTime += deltaTime
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		switch g.currentGameState {
		case MainMenu:

			// TODO: Implement the MainMenu state

			// TODO: Refactor this into a separate function
			rl.ClearBackground(rl.Black)
			rl.DrawText("YAVSC", int32(rl.GetScreenWidth()/2-400), int32(rl.GetScreenHeight()/2-200), 200, rl.Red)
			rl.DrawText("Press Enter to Start", int32(rl.GetScreenWidth()/2-325), int32(rl.GetScreenHeight()/2), 50, rl.White)

			if rl.IsKeyPressed(rl.KeyEnter) {

				// Reset the accumulated time
				accumulatedTime = 0.0

				g.ChangeGameState(Playing)
			}

		case Playing:

			// Update the GameTime in seconds
			g.GameTime += float32(deltaTime)

			rl.BeginMode2D(g.Camera)

			g.Update()

			for accumulatedTime >= float64(g.FixedDeltaTime) {

				// Update Phyisics-related elements
				g.FixedUpdate()

				// Reset the last fixed update time
				// g.LastFixedUpdate = now
				accumulatedTime -= float64(g.FixedDeltaTime)

			}

			// Calculate the interpolation factor for smoother rendering
			interpolation := accumulatedTime / float64(g.FixedDeltaTime)

			// Render the game with interpolation
			g.Render(interpolation)

			rl.EndMode2D()

			// HUD Stuffs Here
			// TODO: Separate the HUD from the PlayerCharacter

			g.Player.HUD.Render(g)
			g.RenderMobsCounter()
			g.RenderPowerUpHUD()

			g.RenderFPS()
		case Paused:
			// TODO: Implement the Paused state

			g.IsPaused = true

		case LeveledUp:
			// Update the music stream
			rl.UpdateMusicStream(SoundTrack)

			// Pause the game
			g.IsPaused = true

			rl.ClearBackground(rl.Black)

			rl.DrawText("Random reward selected", int32(rl.GetScreenWidth()/2-550), int32(rl.GetScreenHeight()/2-200), 100, rl.Yellow)
			rl.DrawText("Press Enter to Continue", int32(rl.GetScreenWidth()/2-300), int32(rl.GetScreenHeight()/2), 50, rl.White)

			if rl.IsKeyPressed(rl.KeyEnter) {

				g.SkillManager.SelectRandomSkill()

				// Reset the accumulated time
				accumulatedTime = 0.0

				g.ChangeGameState(Playing)
			}

		case GameOver:
			// TODO: Implement the GameOver state
			rl.ClearBackground(rl.Black)
			rl.DrawText("Game Over", int32(rl.GetScreenWidth()/2-550), int32(rl.GetScreenHeight()/2-200), 200, rl.Red)
			rl.DrawText("Press Escape to Exit", int32(rl.GetScreenWidth()/2-300), int32(rl.GetScreenHeight()/2), 50, rl.White)

		case Victory:

			victoryTimeMinutes := int(g.GameTime) / 60
			victoryTimeSeconds := int(g.GameTime) % 60
			victoryTimeText := fmt.Sprintf("Time: %d:%d", victoryTimeMinutes, victoryTimeSeconds)

			rl.ClearBackground(rl.Black)
			rl.DrawText("Victory", int32(rl.GetScreenWidth()/2-500), int32(rl.GetScreenHeight()/2-200), 200, rl.Green)
			rl.DrawText(victoryTimeText, int32(rl.GetScreenWidth()/2-200), int32(rl.GetScreenHeight()/2), 50, rl.White)

			rl.DrawText("Press Escape to Exit", int32(rl.GetScreenWidth()/2-300), int32(rl.GetScreenHeight()/2+200), 50, rl.White)

		}

		rl.EndDrawing()
	}

	rl.CloseAudioDevice()
	rl.CloseWindow()

}

func (g *Game) ChangeGameState(state GameState) {

	switch state {
	case MainMenu:
		rl.TraceLog(rl.LogDebug, "Changing GameState to MainMenu")
		g.IsPaused = true
	case Playing:
		rl.TraceLog(rl.LogDebug, "Changing GameState to Playing")
		g.IsPaused = false
	case Paused:
		rl.TraceLog(rl.LogDebug, "Changing GameState to Paused")
		g.IsPaused = true
	case LeveledUp:
		rl.TraceLog(rl.LogDebug, "Changing GameState to LeveledUp")
		g.IsPaused = true
	case GameOver:
		rl.TraceLog(rl.LogDebug, "Changing GameState to GameOver")
		g.IsPaused = true
	case Victory:
		rl.TraceLog(rl.LogDebug, "Changing GameState to Victory")
		g.IsPaused = true
	}

	g.currentGameState = state
}

func (g *Game) SpawnPlayer() {
	g.Player = NewPlayer("Bill", 32, 64, 300, 100, 50)
}

func (g *Game) SpawnProjectile(t ProjectileType, x, y, radius, speed float32, direction rl.Vector2, color rl.Color) {
	g.Projectiles = append(g.Projectiles, NewProjectile(t, x, y, radius, speed, direction, color))
}

func (g *Game) DestroyProjectiles() {

	if len(g.Projectiles) == 0 {
		return
	}

	i := 0

	for _, p := range g.Projectiles {
		if p.Active {
			g.Projectiles[i] = p
			i++
		}
	}

	// Truncate the slice to remove inactive projectiles
	g.Projectiles = g.Projectiles[:i]
}

func (g *Game) SpawnBat() {
	g.Enemies = append(g.Enemies, NewEnemy(TextureAtlas, "Bat", 32, 32, 10))
}

func (g *Game) SpawnEnemy(et EnemyType) {
	switch et {
	case Bat:
		g.SpawnBat()
	case Pumpkin:
		g.SpawnPumpkin()
	}
}

func (g *Game) SpawnPumpkin() {
	g.Enemies = append(g.Enemies, NewEnemy(TextureAtlas, "Pumpkin", 32, 32, 20))
}

func (g *Game) DestroyEnemy() {

	if len(g.Enemies) == 0 {
		return
	}

	i := 0

	for _, e := range g.Enemies {
		if !e.IsDead {
			g.Enemies[i] = e
			i++
		} else {
			g.SpawnExperienceGem(e.X, e.Y, e.ExperienceValue)
		}
	}

	// Truncate the slice to remove dead enemies
	g.Enemies = g.Enemies[:i]

}

func (g *Game) SpawnPowerUp() {

	// Build a new powerup
	powerUp := NewPowerUp()

	// Randomize the PowerUpType
	powerUp.RandomizePowerUpType()

	// Randomize the spawn point
	powerUp.Position = powerUp.RandomizeSpawnPoint()

	// Append the new powerup to the game's powerups
	g.PowerUps = append(g.PowerUps, powerUp)

}

// TODO: Replace temporary function to Render PowerUp HUD
func (g *Game) RenderPowerUpHUD() {
	for i, pu := range g.Player.PowerUps {
		pu.RenderHUD(20 * int32(i))
	}
}

func (g *Game) RenderMobsCounter() {
	mobCount := fmt.Sprintf("Mobs: %d", len(g.Enemies))
	rl.DrawText(mobCount, 10, int32(rl.GetScreenHeight()-50), 20, rl.Red)
}

func (g *Game) RenderFPS() {
	rl.DrawFPS(10, int32(rl.GetScreenHeight()-20))
}

func (g *Game) DestroyPowerUp() {

	if len(g.PowerUps) == 0 {
		return
	}

	// Remove the powerup from the game if its is not active and has been picked up
	i := 0

	for _, pu := range g.PowerUps {
		if !pu.Expired {
			g.PowerUps[i] = pu
			i++
		}
	}

	// Truncate the slice to remove inactive powerups
	g.PowerUps = g.PowerUps[:i]

}

func (g *Game) SpawnExperienceGem(x float32, y float32, amount int32) {
	g.ExperienceGems = append(g.ExperienceGems, NewExperienceGem(x, y, amount))
}

func (g *Game) CreateAllSkills() {

	// Check that we have a valid SkillManager
	if g.SkillManager == nil {
		rl.TraceLog(rl.LogError, "SkillManager is nil")
		return
	}

	// Create a bullet skill
	bullet := NewSkill("Bullet", 10, 200, 1, 8, 1, 0, 10, 400, 0, []UpgradeEffect{
		{CooldownReduction: 0.1},
		{CooldownReduction: 0.2},
		{CooldownReduction: 0.3},
		{CooldownReduction: 0.4},
		{CooldownReduction: 0.5},
		{CooldownReduction: 0.6},
		{CooldownReduction: 0.7},
		{CooldownReduction: 0.8},
	})

	// Create a fireball skill
	fireball := NewSkill("Fireball", 25, 100, 1, 8, 1, 2, 10, 250, 1, []UpgradeEffect{

		{AdditionalDamage: 1, DamageMultiplier: 1, RangeMultiplier: 1, CooldownReduction: 0, IsPiercing: false},
		{AdditionalProjectiles: 1, AdditionalDamage: 25},
		{AdditionalProjectiles: 2},
		{AdditionalProjectiles: 2, DamageMultiplier: 1.5, RangeMultiplier: 1.5},
		{AdditionalProjectiles: 2},
		{AdditionalProjectiles: 3, IsPiercing: true},
		{AdditionalProjectiles: 3},
		{AdditionalProjectiles: 3, DamageMultiplier: 2, RangeMultiplier: 2},
	})

	// Create a magic orb skill
	magicOrb := NewSkill("Magic Orb", 50, 100, 10, 8, 1, 1, 10, 200, 2, []UpgradeEffect{
		{AdditionalDamage: 50},
		{AdditionalProjectiles: 2, AdditionalDamage: 50},
		{AdditionalProjectiles: 3, AdditionalDamage: 50},
		{AdditionalProjectiles: 4, AdditionalDamage: 50},
		{AdditionalProjectiles: 5, AdditionalDamage: 50},
		{AdditionalProjectiles: 6, AdditionalDamage: 50},
		{AdditionalProjectiles: 7, AdditionalDamage: 50},
		{AdditionalProjectiles: 8, AdditionalDamage: 50},
	})

	// Add the skills to the SkillManager
	g.SkillManager.AddSkill(magicOrb)
	g.SkillManager.AddSkill(bullet)
	g.SkillManager.AddSkill(fireball)

	// Add bullet to the active skills
	g.SkillManager.SelectSkill(bullet)
}
