package main

type Entity interface {
	TakeDamage(damage int32)
	GetPosition() (float32, float32)
}
