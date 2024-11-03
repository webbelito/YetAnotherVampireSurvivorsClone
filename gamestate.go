package main

type GameState int

const (
	MainMenu GameState = iota
	Reset
	Playing
	Paused
	LeveledUp
	GameOver
	Victory
)

func (g GameState) String() string {
	return [...]string{"MainMenu", "Reset", "Playing", "Paused", "LeveledUp", "GameOver"}[g]
}
