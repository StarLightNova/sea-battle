package game

import (
    "fmt"
    "github.com/StarLightNova/sea-battle/bin/packages/playersinit"
)

func initializePlayers() playersinit.Players {
    fmt.Println("Initializing Players")

    players, error := playersinit.New()

    if error == nil {
        return players
    }

    fmt.Println("Error! Players initializer did not work.")
    return playersinit.Players{}
}
