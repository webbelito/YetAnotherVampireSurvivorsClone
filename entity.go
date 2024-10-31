package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity interface {
	TakeDamage(damage float32)
	Heal(amount float32)
	GetPosition() (float32, float32)
	GetName() string
	GetCollider() rl.Rectangle
}
