package game

import (
	"sync"
	"time"
)

type Game struct {
    Snake     []Position
    Food      Position
    Width     int
    Height    int
    Direction string
    Mutex     sync.Mutex
}

type Position struct {
    X int
    Y int
}

var GameInstance = Game{
    Snake:     []Position{{10, 10}},
    Food:      Position{5, 5},
    Width:     20,
    Height:    20,
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

func moveSnake() {
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
        resetGame()
        return
    }

    for _, pos := range GameInstance.Snake {
        if pos == newHead {
            resetGame()
            return
        }
    }

    GameInstance.Snake = append([]Position{newHead}, GameInstance.Snake...)

    if newHead == GameInstance.Food {
        GameInstance.Food = Position{(GameInstance.Food.X + 3) % GameInstance.Width, (GameInstance.Food.Y + 3) % GameInstance.Height}
    } else {
        GameInstance.Snake = GameInstance.Snake[:len(GameInstance.Snake)-1]
    }
}

func resetGame() {
    GameInstance.Snake = []Position{{10, 10}}
    GameInstance.Direction = "right"
}
