package main

import (
	"fmt"
	"net/http"
	"snek/server/game"
	"snek/server/handlers"
)

func main() {
    http.HandleFunc("/", handlers.IndexHandler)
    http.HandleFunc("/move", handlers.MoveHandler)

    go game.StartGameLoop()

    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}
