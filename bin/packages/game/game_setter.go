package game

import (
    "fmt"
    "github.com/StarLightNova/sea-battle/bin/packages/playersinit"
)

const (
    FirstPlayer int = 1
    SecondPlayer    = 2
)

func initializePlayers() playersinit.Players {
    fmt.Println("New game created.")
    fmt.Println("Initializing Players...")

    players, error := playersinit.New()

    if error == nil {
        return players
    }

    panic("Error! Players initializer did not work.")
}

func isGameEnded(players playersinit.Players) bool {
    return players.IsOnOfThePlayersDefeated()
}
