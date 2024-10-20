package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Name       string
	X          float32
	Y          float32
	Width      int32
	Height     int32
	Speed      float32
	directionX int32
	directionY int32
}

func NewPlayer(n string, w int32, h int32, s float32) *Player {
	return &Player{
		Name:       n,
		X:          0,
		Y:          0,
		Width:      w,
		Height:     h,
		Speed:      s,
		directionX: 0,
		directionY: 0,
	}
}

func (p *Player) Update() {
	p.HandleInput()
	p.HandleMovment()
	p.Render()
}

func (p *Player) HandleMovment() {
	p.X += float32(p.directionX) * p.Speed * rl.GetFrameTime()
	p.Y += float32(p.directionY) * p.Speed * rl.GetFrameTime()
}

func (p *Player) HandleInput() {
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.directionX = 1
	} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.directionX = -1
	} else {
		p.directionX = 0
	}

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.directionY = -1
	} else if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.directionY = 1
	} else {
		p.directionY = 0
	}
}

func (p *Player) Render() {
	rl.DrawRectangle(int32(p.X), int32(p.Y), p.Width, p.Height, rl.Green)
}

func (p *Player) GetName() string {
	return p.Name
}
