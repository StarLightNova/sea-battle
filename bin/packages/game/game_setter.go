package game

import (
    "fmt"
    "github.com/StarLightNova/sea-battle/bin/packages/playersinit"
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
