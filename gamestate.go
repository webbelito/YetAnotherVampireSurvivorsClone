package main

type GameState int

const (
	MainMenu GameState = iota
	Reset
	Playing
	Paused
	LeveledUp
	GameOver
)
