package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Tile Types
const (
	Empty = iota
	Pillar
	LeftWall
	RightWall
	TopWall
	BottomWall
	BlockWall
	Grass
)

// Map of characters to tile types
var tileMap = map[rune]int{
	' ': Empty,
	'X': Pillar,
	'[': LeftWall,
	']': RightWall,
	'^': TopWall,
	'_': BottomWall,
	'@': BlockWall,
	'.': Grass,
}

// Level represents a game level with a grid of tiles
type Level struct {
	Width  int
	Height int
	Tiles  [][]int
	Walls  []Wall
}

type Wall struct {
	Collider rl.Rectangle
}

// Create a new level based on the provided string grid
func NewLevel(grid []string) *Level {

	height := len(grid)
	width := len(grid[0])
	tiles := make([][]int, height)
	walls := make([]Wall, 0)

	for i := 0; i < height; i++ {

		// Create a new row of tiles
		tiles[i] = make([]int, width)

		for j, char := range grid[i] {

			if tileType, ok := tileMap[char]; ok {
				tiles[i][j] = tileType

				// If the tile is a wall, create a new wall object
				switch tileType {
				case Pillar, LeftWall, RightWall, TopWall, BottomWall, BlockWall:
					// Create a new wall
					wall := Wall{
						Collider: rl.NewRectangle(float32(j*32), float32(i*32), 32, 32),
					}
					// Append the wall to the level's walls
					walls = append(walls, wall)
				}
			} else {
				rl.TraceLog(rl.LogError, "Unknown tile type: %c", char)
			}
		}
	}

	return &Level{
		Width:  width,
		Height: height,
		Tiles:  tiles,
		Walls:  walls,
	}
}

func (l *Level) Render() {
	for i := 0; i < l.Height; i++ {
		for j := 0; j < l.Width; j++ {

			x := int32(j * 32)
			y := int32(i * 32)

			switch l.Tiles[i][j] {
			case Pillar:

				rl.DrawTexturePro(
					TextureAtlas,
					rl.NewRectangle(128, 320, 32, 32),
					rl.NewRectangle(float32(x), float32(y), 32, 32),
					rl.NewVector2(0, 0),
					0,
					rl.White,
				)

			case LeftWall:

				rl.DrawTexturePro(
					TextureAtlas,
					rl.NewRectangle(96, 320, 32, 32),
					rl.NewRectangle(float32(x), float32(y), 32, 32),
					rl.NewVector2(0, 0),
					0,
					rl.White,
				)

			case RightWall:

				rl.DrawTexturePro(
					TextureAtlas,
					rl.NewRectangle(32, 320, 32, 32),
					rl.NewRectangle(float32(x), float32(y), 32, 32),
					rl.NewVector2(0, 0),
					0,
					rl.White,
				)

			case TopWall:

				rl.DrawTexturePro(
					TextureAtlas,
					rl.NewRectangle(0, 320, 32, 32),
					rl.NewRectangle(float32(x), float32(y), 32, 32),
					rl.NewVector2(0, 0),
					0,
					rl.White,
				)

			case BottomWall:

				rl.DrawTexturePro(
					TextureAtlas,
					rl.NewRectangle(64, 320, 32, 32),
					rl.NewRectangle(float32(x), float32(y), 32, 32),
					rl.NewVector2(0, 0),
					0,
					rl.White,
				)

			case BlockWall:

				rl.DrawTexturePro(
					TextureAtlas,
					rl.NewRectangle(160, 320, 32, 32),
					rl.NewRectangle(float32(x), float32(y), 32, 32),
					rl.NewVector2(0, 0),
					0,
					rl.White,
				)

			case Grass:

				rl.DrawTexturePro(
					TextureAtlas,
					rl.NewRectangle(0, 352, 32, 32),
					rl.NewRectangle(float32(x), float32(y), 32, 32),
					rl.NewVector2(0, 0),
					0,
					rl.White,
				)
			case Empty:
				rl.TraceLog(rl.LogError, "Empty tile at (%d, %d)", i, j)
			default:
				rl.TraceLog(rl.LogError, "Unknown tile type: %d", l.Tiles[i][j])
			}
		}
	}
}

func (l *Level) CheckCollisions(e Entity) bool {
	for _, wall := range l.Walls {
		if rl.CheckCollisionRecs(wall.Collider, e.GetCollider()) {
			return true
		}
	}

	return false
}
