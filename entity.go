package main

type Entity interface {
	TakeDamage(damage float32)
	Heal(amount float32)
	GetPosition() (float32, float32)
	GetName() string
}
