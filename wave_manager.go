package main

type EnemyType int

const (
	Bat EnemyType = iota
	Pumpkin
)

type Wave struct {
	EnemyCounts   map[EnemyType]int
	SpawnInterval float32
	Duration      float32
	SpawnTimer    float32
}

type WaveManager struct {
	Waves              []Wave
	CurrentWaveIndex   int
	TimeSinceLastWave  float32
	TimeSinceLastSpawn float32
}

func (wm *WaveManager) SpawnEnemies(g *Game) {
	currentWave := wm.Waves[wm.CurrentWaveIndex]

	if wm.TimeSinceLastSpawn >= currentWave.SpawnInterval {
		for enemyType, count := range currentWave.EnemyCounts {
			if count > 0 {
				g.SpawnEnemy(enemyType)
				wm.Waves[wm.CurrentWaveIndex].EnemyCounts[enemyType]--
			}
		}

		wm.TimeSinceLastSpawn = 0
	}
}
