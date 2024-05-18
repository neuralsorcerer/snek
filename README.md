# snek

A simple Snake game implemented using Go + HTMX. The game is served via an HTTP server and can be played in a web browser.

## Features

- The snake moves automatically and grows when it eats food.
- The snake cannot move beyond the boundaries or hit itself.
- Simple web interface using htmx for interactions.


## Getting Started

### Clone the Repository

```sh
git clone https://github.com/your-username/snake-game.git
cd snake-game
```

### Initialize the Go Module

```sh
go mod init snek
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

### Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
