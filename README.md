# snek

A simple Snake game implemented using Go + HTMX. The game is served via an HTTP server and can be played in a web browser.

## Features

- The snake moves automatically and grows when it eats food.
- The snake cannot move beyond the boundaries or hit itself.
- Simple web interface using htmx for interactions.
- The current score and high score are shown and update as you play.
- Game over message with a reset button to start a new round.
- Supports both arrow keys and WASD for controlling the snake.

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/neuralsorcerer/snek.git
cd snek
```

### Initialize the Go Module

```sh
go mod tidy
```

### Run the Server

```sh
go run server/main.go
```

The server will start at `http://localhost:8080`.

### Play the Game

Open your web browser and navigate to `http://localhost:8080` to start playing the game.

### How to Play

- Use the arrow keys to change the direction of the snake.
- Avoid running into the walls or the snake's own body.
- Use the arrow keys or WASD to change the direction of the snake.
- Avoid running into the walls or the snake's own body. Press **Reset** or the `r` key to start a new round after a game over.

### Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
