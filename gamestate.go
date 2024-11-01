package main

type GameState int

const (
	MainMenu GameState = iota
	Playing
	Paused
	LeveledUp
	GameOver
)
