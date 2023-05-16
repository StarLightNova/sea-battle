package playersinit

import (
    "fmt"
    "github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

func (players Players) EntireMap() string {
    entireMap := fmt.Sprintf("%s\t\t%s\n", playermap.YOUR_BOARD_NAME, playermap.OPPONENT_BOARD_NAME)

    entireMap += fmt.Sprintf("%s\t\t%s\n", playermap.MAP_COLUMNS, playermap.MAP_COLUMNS)
    entireMap += fmt.Sprintf("%s\t\t%s\n", playermap.COLUMN_BOTTOM_DIVIDER, playermap.COLUMN_BOTTOM_DIVIDER)

    for _, key := range playermap.GetLetterCoordinates() {
        entireMap += fmt.Sprintf(" %s )", key)

        for _, letter := range players.FirstPlayer.PlayerMap.TheMap[key] {
            entireMap += fmt.Sprintf(" %s |", letter)
        }

        entireMap += fmt.Sprint("\t\t")

        entireMap += fmt.Sprintf(" %s )", key)

        for _, letter := range players.SecondPlayer.PlayerMap.ShadowMap[key] {
            entireMap += fmt.Sprintf(" %s |", letter)
        }

        entireMap += fmt.Sprintf("\n%s\t\t%s\n", playermap.BOTTOM_DIVIDER, playermap.BOTTOM_DIVIDER)
    }

    return entireMap
}

