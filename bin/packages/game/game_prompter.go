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

        if loweredUsersPrompt == YES {
            return true
        } else if loweredUsersPrompt == NO {
            return false
        } else {
            fmt.Println("Your reply is not understandable by computer. (y/n)")
        }
    }
}

func makeMove(attacker, opponent playermap.PlayerMap) {
    var userInput string

    howToAttackCoordinates()

    fmt.Scanln(&userInput)

    opponent.GetDamage(coordinates.StringToRowColumn(userInput))
}

func botMakeMove(attacker, opponent playermap.PlayerMap) {
    opponent.GetDamage(coordinates.RandomHalfCoordinates())
}
