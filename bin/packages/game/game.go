package game

import (
    "github.com/StarLightNova/sea-battle/bin/packages/playersinit"
)

func New() {
    players := initializePlayers()
    players.PutShips()

    announceAGame()
    yourBoard(players.FirstPlayer.PlayerMap)

    readyToGo()

    if promptToStartAGame() {
        startGame(players)
    }

    goodbye()
}

func startGame(players playersinit.Players) {
    // TODO
}
