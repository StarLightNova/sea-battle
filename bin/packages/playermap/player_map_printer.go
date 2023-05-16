package playermap

import "fmt"

func mapPrinter(playerMap map[string][]string) string {
    singleMap := fmt.Sprintf("%s\n", MAP_COLUMNS)

    singleMap += fmt.Sprintf("%s\n", COLUMN_BOTTOM_DIVIDER)

    for _, key := range GetLetterCoordinates() {
        singleMap += fmt.Sprintf(" %s )", key)

        for _, letter := range playerMap[key] {
            singleMap += fmt.Sprintf(" %s |", letter)
        }

        singleMap += fmt.Sprintf("\n%s\n", BOTTOM_DIVIDER)
    }

    return singleMap
}
