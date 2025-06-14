package game

import (
	"math/rand"
	"sync"
	"time"
)

type Game struct {
	Snake     []Position
	Food      Position
	Width     int
	Height    int
	Score     int
	HighScore int
	GameOver  bool
	Direction string
	Mutex     sync.Mutex
}

type Position struct {
	X int
	Y int
}

var GameInstance = Game{
	Snake:     []Position{{10, 10}},
	Width:     20,
	Height:    20,
	Score:     0,
	HighScore: 0,
	GameOver:  false,
	Direction: "right",
}

func StartGameLoop() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		GameInstance.Mutex.Lock()

		moveSnake()

		GameInstance.Mutex.Unlock()
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	GameInstance.Food = generateFoodPosition()
}

func generateFoodPosition() Position {
	for {
		p := Position{rand.Intn(GameInstance.Width), rand.Intn(GameInstance.Height)}
		conflict := false
		for _, pos := range GameInstance.Snake {
			if pos == p {
				conflict = true
				break
			}
		}
		if !conflict {
			return p
		}
	}
}

func moveSnake() {
	if GameInstance.GameOver {
		return
	}

	head := GameInstance.Snake[0]
	var newHead Position

	switch GameInstance.Direction {
	case "up":
		newHead = Position{head.X, head.Y - 1}
	case "down":
		newHead = Position{head.X, head.Y + 1}
	case "left":
		newHead = Position{head.X - 1, head.Y}
	case "right":
		newHead = Position{head.X + 1, head.Y}
	}

	if newHead.X < 0 || newHead.X >= GameInstance.Width || newHead.Y < 0 || newHead.Y >= GameInstance.Height {
		endGame()
		return
	}

	for _, pos := range GameInstance.Snake {
		if pos == newHead {
			endGame()
			return
		}
	}

	GameInstance.Snake = append([]Position{newHead}, GameInstance.Snake...)

	if newHead == GameInstance.Food {
		GameInstance.Score++
		GameInstance.Food = generateFoodPosition()
	} else {
		GameInstance.Snake = GameInstance.Snake[:len(GameInstance.Snake)-1]
	}
}

func endGame() {
	if GameInstance.Score > GameInstance.HighScore {
		GameInstance.HighScore = GameInstance.Score
	}
	GameInstance.GameOver = true
}

func resetGame() {
	GameInstance.Snake = []Position{{10, 10}}
	GameInstance.Direction = "right"
	GameInstance.Score = 0
	GameInstance.Food = generateFoodPosition()
	GameInstance.GameOver = false
}

func ResetGame() {
	GameInstance.Mutex.Lock()
	defer GameInstance.Mutex.Unlock()
	if GameInstance.Score > GameInstance.HighScore {
		GameInstance.HighScore = GameInstance.Score
	}
	resetGame()
}