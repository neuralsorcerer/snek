package main

import (
	"fmt"
	"net/http"

	"github.com/neuralsorcerer/snek/server/game"
	"github.com/neuralsorcerer/snek/server/handlers"
)

func main() {
    http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/move", handlers.MoveHandler)
	http.HandleFunc("/score", handlers.ScoreHandler)
	http.HandleFunc("/reset", handlers.ResetHandler)

    go game.StartGameLoop()

    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}
