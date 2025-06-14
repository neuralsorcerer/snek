package handlers

import (
	"fmt"
	"net/http"

	"github.com/neuralsorcerer/snek/server/game"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func MoveHandler(w http.ResponseWriter, r *http.Request) {
    direction := r.URL.Query().Get("direction")

    game.GameInstance.Mutex.Lock()
    defer game.GameInstance.Mutex.Unlock()

    if direction != "" {
        game.GameInstance.Direction = direction
    }

    renderGame(w)
}

func ScoreHandler(w http.ResponseWriter, r *http.Request) {
	game.GameInstance.Mutex.Lock()
	score := game.GameInstance.Score
	high := game.GameInstance.HighScore
	over := game.GameInstance.GameOver
	game.GameInstance.Mutex.Unlock()

	if over {
		fmt.Fprintf(w, "Game Over! Score: %d | Best: %d", score, high)
	} else {
		fmt.Fprintf(w, "Score: %d | Best: %d", score, high)
	}
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	game.ResetGame()

	game.GameInstance.Mutex.Lock()
	defer game.GameInstance.Mutex.Unlock()
	renderGame(w)
}
func renderGame(w http.ResponseWriter) {
    var gridHTML string
	for y := 0; y < game.GameInstance.Height; y++ {
		for x := 0; x < game.GameInstance.Width; x++ {
			cellClass := "cell"
			for _, pos := range game.GameInstance.Snake {
				if pos.X == x && pos.Y == y {
					cellClass = "cell snake"
					break
				}
			}
			if game.GameInstance.Food.X == x && game.GameInstance.Food.Y == y {
				cellClass = "cell food"
			}
			gridHTML += fmt.Sprintf(`<div class="%s"></div>`, cellClass)
		}
	}
	fmt.Fprint(w, gridHTML)
}
