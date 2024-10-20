package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Name   string
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func NewPlayer(n string, w int32, h int32) *Player {
	return &Player{
		Name:   n,
		X:      0,
		Y:      0,
		Width:  w,
		Height: h,
	}
}

func (p *Player) Render() {
	rl.DrawRectangle(p.X, p.Y, p.Width, p.Height, rl.Green)
}

func (p *Player) GetName() string {
	return p.Name
}
