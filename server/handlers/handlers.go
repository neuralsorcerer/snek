package handlers

import (
	"fmt"
	"net/http"
	"snek/server/game"
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
