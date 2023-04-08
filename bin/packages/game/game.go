package game

import (
    "fmt"
)

func New() {
    fmt.Println("New game created")

    players := initializePlayers()
    players.PutShips()
}
