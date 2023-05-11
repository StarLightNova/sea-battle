package game

import (
    "fmt"
    "strings"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
	"github.com/StarLightNova/sea-battle/bin/packages/coordinates"
)	

func promptToStartAGame() bool {
    var usersInput string

    for {
        fmt.Scanln(&usersInput)

        loweredUsersPrompt := strings.ToLower(usersInput)

        if loweredUsersPrompt == "y" {
            return true
        } else if loweredUsersPrompt == "n" {
            return false
        } else {
            fmt.Println("Your reply is not understandable by computer. (y/n)")
        }
    }
}

func makeMove(attacker, opponent playermap.PlayerMap) {
    opponentsShadowMap(opponent)

    var userInput string

    howToAttackCoordinates()

    fmt.Scanln(&userInput)

    opponent.GetDamage(coordinates.StringToRowColumn(userInput))
    
    opponentsShadowMap(opponent)
}

func botMakeMove(attacker, opponent playermap.PlayerMap) {
    opponent.GetDamage(coordinates.RandomHalfCoordinates())

    yourBoardHasAttacked(opponent)
}
