package game

import (
    "fmt"
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
    clearConsole()
    fmt.Println(players.EntireMap())

    turn := FirstPlayer
    
    for !isGameEnded(players) {
        
        switch turn {
        case FirstPlayer:
            makeMove(players.FirstPlayer.PlayerMap, players.SecondPlayer.PlayerMap)

            turn = SecondPlayer
        case SecondPlayer:
            botMakeMove(players.SecondPlayer.PlayerMap, players.FirstPlayer.PlayerMap)

            turn = FirstPlayer
        }

        clearConsole()
        fmt.Println(players.EntireMap())
    }
}

func isGameEnded(players playersinit.Players) bool {
    return players.IsOnOfThePlayersDefeated()
}
