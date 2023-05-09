package game

import (
    "fmt"
    "strings"
    "regexp"
    "strconv"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
    "math/rand"
	"time"

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

func formatCoordinates(userCoordindates string) (row string, column int) {
    regexPattern, _ := regexp.Compile("(\\w{1})(10|\\d{1})")

    if regexPattern.MatchString(userCoordindates) {
        matchedGroupes := regexPattern.FindStringSubmatch(userCoordindates)
        row = strings.ToUpper(matchedGroupes[1])
        column, _ = strconv.Atoi(matchedGroupes[2])
    }

    return
}

func generateRandomCoordinate() (row string, column int) {
    rand.Seed(time.Now().UnixNano())
    row = playermap.GetLetterCoordinates()[rand.Intn(len(playermap.GetLetterCoordinates()))]
    column = rand.Intn(10) + 1

    return
}

func makeMove(attacker, opponent playermap.PlayerMap) {
    opponentsShadowMap(opponent)

    var userInput string

    howToAttackCoordinates()

    fmt.Scanln(&userInput)

    opponent.GetDamage(formatCoordinates(userInput))
    
    opponentsShadowMap(opponent)
}

func botMakeMove(attacker, opponent playermap.PlayerMap) {
    opponent.GetDamage(generateRandomCoordinate())
    yourBoardHasAttacked(opponent)
}
