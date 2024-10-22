package game

import (
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/enemy"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/player"
	"github.com/webbelito/YetAnotherVampireSurvivorsClone/projectile"

	// 3rd-party library
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Player      *player.PlayerCharacter
	Enemies     []*enemy.Enemy
	Projectiles []*projectile.Projectile
}

func NewGame() *Game {
	return &Game{
		Enemies:     make([]*enemy.Enemy, 0),
		Projectiles: make([]*projectile.Projectile, 0),
	}
}

func (g *Game) Update() {

	if g.Player == nil {
		g.SpawnPlayer()
	}

	// Spawn an enemy
	if rl.IsKeyPressed(rl.KeySpace) {
		g.SpawnEnemy()
	}

	g.Player.Update()

	for i := 0; i < len(g.Enemies); i++ {
		g.Enemies[i].Update(g.Player)
	}

	for i := 0; i < len(g.Projectiles); i++ {
		g.Projectiles[i].Update()
	}

}

func (g *Game) SpawnPlayer() {
	g.Player = player.NewPlayer("Player", 50, 100, 150, 100, 50)
}

func (g *Game) SpawnProjectile(x, y, radius, speed float32, direction rl.Vector2) {
	g.Projectiles = append(g.Projectiles, projectile.NewProjectile(x, y, radius, speed, direction))
}

func (g *Game) SpawnEnemy() {
	g.Enemies = append(g.Enemies, enemy.NewEnemy("BAT", 50, 50, 100, 100, 50))
}
