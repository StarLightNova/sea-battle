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
    clearConsole()

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
    }
}

